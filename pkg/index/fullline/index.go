package fullline

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/zu1k/she/persistence"

	"github.com/zu1k/she/internal/processor"

	C "github.com/zu1k/she/constant"
	"github.com/zu1k/she/pkg/source"

	"github.com/blevesearch/bleve"
	"github.com/cheggaaa/pb/v3"
	"github.com/zu1k/she/pkg/tools"
)

func newBleveIndex(blevepath string) (index bleve.Index, err error) {
	mapping := bleve.NewIndexMapping()
	err = mapping.AddCustomTokenizer("sego",
		map[string]interface{}{
			"dictpath": "dictionary.txt",
			"type":     "sego",
		},
	)
	if err != nil {
		panic(err)
	}
	err = mapping.AddCustomAnalyzer("sego",
		map[string]interface{}{
			"type":      "sego",
			"tokenizer": "sego",
		},
	)
	if err != nil {
		panic(err)
	}
	mapping.DefaultAnalyzer = "sego"
	index, err = bleve.NewUsing(blevepath, mapping, "scorch", "scorch", map[string]interface{}{
		"forceSegmentType":    "zap",
		"forceSegmentVersion": 12,
	})
	return
}

func ParseAndIndex(filePath string) {
	fi, err := os.Open(filePath)
	if err != nil {
		for {
			time.Sleep(time.Second)
			fi, err = os.Open(filePath)
			if err == nil {
				break
			}
		}
	}

	lineNum, err := tools.LineCounter(filePath)
	if err != nil {
		panic(err)
	}

	br := bufio.NewReader(fi)

	infoChan := make(chan string, 1000)
	fileName := tools.Path2Name(filePath)
	dirPath := tools.Path2Path(filePath)
	storePath := filepath.Join(C.Path.IndexDir(), dirPath, fileName)
	fmt.Println("index store path", storePath, dirPath)

	os.RemoveAll(storePath)
	indexer, err := newBleveIndex(storePath)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				close(infoChan)
				break
			}
			infoChan <- string(a)
		}
	}()

	indexProcessor(indexer, infoChan, lineNum)
	processor.AddSource(fileName, source.BleveIndex, storePath)
	_ = persistence.NewSource(fileName, source.BleveIndex, storePath, filePath)
}

func indexProcessor(index bleve.Index, infoChan chan string, lineCount int) {
	batch := index.NewBatch()
	//linenum := 0
	//for line := range infoChan {
	//	linenum++
	//	fmt.Println(line)
	//	_ = batch.Index(fmt.Sprintf("%d", linenum), line)
	//}

	linenum := 0
	i := 0

	bar := pb.StartNew(lineCount)
	for line := range infoChan {
		i++
		linenum++
		_ = batch.Index(fmt.Sprintf("%d", linenum), line)
		bar.Increment()
		if i == 5000 {
			_ = index.Batch(batch)
			i = 0
			batch = index.NewBatch()
		}
	}
	if i > 0 {
		_ = index.Batch(batch)
	}
	index.Close()
	time.Sleep(time.Second)
	bar.Finish()
	fmt.Printf("finish: %d valid rows indexed\n", linenum)
}
