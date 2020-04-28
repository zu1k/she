package jiudian2000w

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"

	"github.com/blevesearch/bleve"
	"github.com/zu1k/she/source/bleveidx"
)

func openCSC(filepath string) (csvReader *csv.Reader, err error) {
	cntb, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	csvReader = csv.NewReader(strings.NewReader(string(cntb)))
	return
}

func lineCounter(filepath string) (int, error) {
	r, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	var readSize int
	var count int

	buf := make([]byte, 1024)

	for {
		readSize, err = r.Read(buf)
		if err != nil {
			break
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			if i == -1 || readSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
	}
	if readSize > 0 && count == 0 || count > 0 {
		count++
	}
	if err == io.EOF {
		return count, nil
	}

	return count, err
}

func ParseAndIndex(filepath string) {
	reader, err := openCSC(filepath)
	if err != nil {
		panic(err)
	}

	lineNum, err := lineCounter(filepath)
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
	indexer, err := newBleveIdx(storePath)
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

func newBleveIdx(kufilepath string) (index bleve.Index, err error) {
	index, err = bleveidx.NewBleveIndex(kufilepath, 2)
	if err != nil {
		panic(err)
	}
	return
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
