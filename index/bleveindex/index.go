package bleveindex

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/cheggaaa/pb/v3"
	"github.com/zu1k/she/index/tools"
	"github.com/zu1k/she/source/bleveidx"
)

func ParseAndIndex(filepath, infoFilePath string) {
	lineNum, err := tools.LineCounter(filepath)
	if err != nil {
		panic(err)
	}

	entityChan := make(chan Entity, 1000)
	fileName := path.Base(filepath)
	if runtime.GOOS == "windows" {
		filepaths := strings.Split(filepath, "\\")
		fileName = filepaths[len(filepaths)-1]
	}

	storePath := "D:\\sheku\\" + fileName
	_ = os.RemoveAll(storePath)
	indexer, err := bleveidx.NewBleveIndex(storePath, 2)
	if err != nil {
		panic(err)
	}

	go Parse(filepath, infoFilePath, entityChan)
	indexProcessor(indexer, entityChan, lineNum)
}

func indexProcessor(index bleve.Index, entityChan chan Entity, lineCount int) {
	batch := index.NewBatch()
	linenum := 0
	i := 0
	bar := pb.StartNew(lineCount)
	for entity := range entityChan {
		i++
		linenum++
		_ = batch.Index(fmt.Sprintf("%d", linenum), entity)
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
