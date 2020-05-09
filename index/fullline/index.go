package fullline

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/zu1k/she/persistence"
	"github.com/zu1k/she/source"

	"github.com/blevesearch/bleve"
	"github.com/cheggaaa/pb/v3"
	"github.com/zu1k/she/index/tools"
)

func newBleveIndex(blevepath string) (index bleve.Index, err error) {
	mapping := bleve.NewIndexMapping()
	//==========use seg==============
	err = mapping.AddCustomTokenizer("sego",
		map[string]interface{}{
			"dictpath": "C:\\Users\\zu1k\\go\\pkg\\mod\\github.com\\huichen\\sego@v0.0.0-20180617034105-3f3c8a8cfacc\\data\\dictionary.txt",
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
	//==========use seg==============

	index, err = bleve.NewUsing(blevepath, mapping, "scorch", "scorch", map[string]interface{}{
		"forceSegmentType":    "zap",
		"forceSegmentVersion": 12,
	})
	return
}

func ParseAndIndex(filepath string) {
	fi, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	lineNum, err := tools.LineCounter(filepath)
	if err != nil {
		panic(err)
	}

	br := bufio.NewReader(fi)

	infoChan := make(chan string, 1000)
	fileName := path.Base(filepath)
	if runtime.GOOS == "windows" {
		filepaths := strings.Split(filepath, "\\")
		fileName = filepaths[len(filepaths)-1]
	}

	storePath := "D:\\sheku\\" + fileName
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
	_ = persistence.NewSource(fileName, source.BleveIndex, storePath)
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
	time.Sleep(time.Second)
	bar.Finish()
	fmt.Printf("finish: %d valid rows indexed\n", linenum)
}
