package bleveidx

import (
	"bufio"
	"fmt"
	"github.com/blevesearch/bleve/index/scorch"
	"io"
	"os"

	"github.com/blevesearch/bleve"
	//_ "github.com/sunvim/bleve-sego"
)

func NewBleveIndex(blevepath string, storeType int) (index bleve.Index, err error) {
	if storeType == 2 {
		bleve.Config.DefaultIndexType = scorch.Name
	}
	mapping := bleve.NewIndexMapping()

	////==========use seg==============
	//err = mapping.AddCustomTokenizer("sego",
	//	map[string]interface{}{
	//		"dictpath": "C:\\Users\\zu1k\\go\\pkg\\mod\\github.com\\huichen\\sego@v0.0.0-20180617034105-3f3c8a8cfacc\\data\\dictionary.txt",
	//		"type":     "sego",
	//	},
	//)
	//if err != nil {
	//	panic(err)
	//}
	//err = mapping.AddCustomAnalyzer("sego",
	//	map[string]interface{}{
	//		"type":      "sego",
	//		"tokenizer": "sego",
	//	},
	//)
	//if err != nil {
	//	panic(err)
	//}
	//mapping.DefaultAnalyzer = "sego"
	////==========use seg==============

	index, err = bleve.New(blevepath, mapping)
	return
}

func readFileByRow(path string, lineChan chan string) {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lineChan <- string(a)
	}
	close(lineChan)
}
