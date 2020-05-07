package bleveindex

import (
	"fmt"
	"testing"

	"github.com/blevesearch/bleve/index/scorch"

	"github.com/blevesearch/bleve"
)

//func TestQQGroupSearch(t *testing.T) {
//	b := newBleveIdx("../../ku/test1")
//	results := b.Search("zhaoyanxia002")
//	if len(results) > 0 {
//		for _, i := range results {
//			fmt.Println(i.Text)
//		}
//	} else {
//		t.Errorf("not found")
//	}
//}

func TestBleveTest(t *testing.T) {
	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleveindex indexing is easy",
	}

	bleve.Config.DefaultIndexType = scorch.Name
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("../../ku/bleveindex/test4", mapping)

	if err != nil {
		panic(err)
	}
	_ = index.Index(message.Id, message)
	//index, err := bleveindex.Open("../../ku/bleveindex/test")
	//if err != nil {
	//	panic(err)
	//}
	err = index.Index("test_full_text_idx", "test_full_text example")
	if err != nil {
		panic(err)
	}
	query := bleve.NewQueryStringQuery("example")
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResult, _ := index.Search(searchRequest)
	fmt.Println(searchResult.Hits[0].Fields)
	fmt.Println(searchResult.Hits[1].Fields)
}

func TestNewBleve(t *testing.T) {
	//NewBleve("D:\\Project\\she\\ku\\12306\\account.csv", "D:\\Project\\she\\ku\\12306\\bleveindex")
	//NewBleveScorch("D:\\Project\\she\\ku\\12306\\account.csv", "D:\\Project\\she\\ku\\12306\\bleveindex")
}
