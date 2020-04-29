package jiudian2000w

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

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

	doc := bleve.NewDocumentMapping()

	docName := bleve.NewTextFieldMapping()
	docCtfid := bleve.NewTextFieldMapping()
	docGender := bleve.NewTextFieldMapping()
	docBirthday := bleve.NewTextFieldMapping()
	docAddress := bleve.NewTextFieldMapping()
	docEmail := bleve.NewTextFieldMapping()
	docMobile := bleve.NewTextFieldMapping()

	doc.AddFieldMappingsAt("name", docName)
	doc.AddFieldMappingsAt("ctfid", docCtfid)
	doc.AddFieldMappingsAt("gender", docGender)
	doc.AddFieldMappingsAt("birthday", docBirthday)
	doc.AddFieldMappingsAt("address", docAddress)
	doc.AddFieldMappingsAt("email", docEmail)
	doc.AddFieldMappingsAt("mobile", docMobile)

	mapping.AddDocumentMapping("people", doc)
	index, err = bleve.NewUsing(blevepath, mapping, "scorch", "scorch", map[string]interface{}{
		"forceSegmentType":    "zap",
		"forceSegmentVersion": 12,
	})
	return
}

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
	indexer, err := newBleveIndex(storePath)
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
