package bleveidx

import (
	"fmt"

	"github.com/zu1k/she/source"

	"github.com/blevesearch/bleve"
	"github.com/zu1k/she/log"
)

type bleveidx struct {
	index bleve.Index
}

func init() {
	source.Register("bleveidx", newBleveIdx)
}

func newBleveIdx(info interface{}) source.Source {
	path := info.(string)
	index, err := bleve.Open(path)
	if err != nil {
		log.Errorln("Fail to open bleve index file")
		return nil
	}
	return &bleveidx{index: index}
}

// GetName return bleveidx name
func (b *bleveidx) GetName() string {
	return "BleveIdx"
}

// Search return result slice from source bleveidx
func (b *bleveidx) Search(key interface{}) (results []source.Result) {
	str := key.(string)
	log.Infoln("Search BleveIdx, key = %s", key)
	query := bleve.NewMatchQuery(str)
	search := bleve.NewSearchRequest(query)
	searchResults, err := b.index.Search(search)
	if err != nil {
		return
	}
	//TODO 查找到索引后找真实数据
	hits := searchResults.Hits
	for _, i := range hits {
		fmt.Println(i.Index)
	}
	return
}
