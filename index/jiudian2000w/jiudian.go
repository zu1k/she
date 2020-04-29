package jiudian2000w

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/zu1k/she/source/bleveidx"

	"github.com/zu1k/she/index/tools"

	"github.com/cheggaaa/pb/v3"

	"github.com/blevesearch/bleve"
)

func ParseAndIndex(filepath string) {
	reader, err := tools.OpenCSC(filepath)
	if err != nil {
		panic(err)
	}

	lineNum, err := tools.LineCounter(filepath)
	if err != nil {
		panic(err)
	}

	infoChan := make(chan People, 1000)
	fileName := path.Base(filepath)
	if runtime.GOOS == "windows" {
		filepaths := strings.Split(filepath, "\\")
		fileName = filepaths[len(filepaths)-1]
	}

	storePath := "D:\\sheku\\" + fileName
	os.RemoveAll(storePath)
	indexer, err := bleveidx.NewBleveIndex(storePath, 2)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			record, err := reader.Read()
			if err != nil {
				switch err {
				case io.EOF:
					close(infoChan)
					return
				default:
					continue
				}
			}
			people, err := parsePeopleInfo(record)
			if err != nil {
				continue
			}
			infoChan <- people
		}
	}()

	indexProcessor(indexer, infoChan, lineNum)
}

func indexProcessor(index bleve.Index, infoChan chan People, lineCount int) {
	batch := index.NewBatch()
	//linenum := 0
	//for people := range infoChan {
	//	linenum++
	//	fmt.Println(people)
	//	_ = batch.Index(fmt.Sprintf("%d", linenum), people)
	//}

	linenum := 0
	i := 0

	bar := pb.StartNew(lineCount)
	for people := range infoChan {
		i++
		linenum++
		_ = batch.Index(fmt.Sprintf("%d", linenum), people)
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
