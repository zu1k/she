package bleveindex

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/cheggaaa/pb/v3"
	C "github.com/zu1k/she/constant"
	"github.com/zu1k/she/persistence"
	"github.com/zu1k/she/pkg/source"
	"github.com/zu1k/she/pkg/source/bleveindex"
	"github.com/zu1k/she/pkg/tools"
)

func ParseAndIndex(filePath, infoFilePath string) {
	lineNum, err := tools.LineCounter(filePath)
	if err != nil {
		panic(err)
	}

	entityChan := make(chan Entity, 1000)

	fileName := tools.Path2Name(filePath)
	storePath := filepath.Join(C.Path.IndexDir(), fileName)
	_ = os.RemoveAll(storePath)
	indexer, err := bleveindex.NewBleveIndex(storePath, 2)
	if err != nil {
		panic(err)
	}

	go Parse(filePath, infoFilePath, entityChan)
	indexProcessor(indexer, entityChan, lineNum)
	_ = persistence.NewSource(fileName, source.BleveIndex, storePath, filePath)
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
