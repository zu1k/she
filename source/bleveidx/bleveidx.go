package bleveidx

import (
	"fmt"
	"sync"

	"github.com/zu1k/she/common"
	"github.com/zu1k/she/source"

	"github.com/blevesearch/bleve"
	"github.com/zu1k/she/log"
)

type Bleveidx struct {
	index bleve.Index
}

func init() {
	source.Register("Bleveidx", OpenBleveIdx)
}

func OpenBleveIdx(info interface{}) source.Source {
	path := info.(string)
	index, err := bleve.Open(path)
	if err != nil {
		log.Errorln("Fail to open bleve index file")
		return nil
	}
	return &Bleveidx{index: index}
}

// GetName return Bleveidx name
func (b *Bleveidx) GetName() string {
	return "BleveIdx"
}

// Search return result slice from source Bleveidx
func (b *Bleveidx) Search(key interface{}, resChan chan common.Result, wg *sync.WaitGroup) {
	str := key.(string)
	log.Infoln("Search BleveIdx, key = %s", key)
	query := bleve.NewMatchQuery(str)
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResults, err := b.index.Search(searchRequest)
	if err != nil {
		return
	}
	//TODO 查找到索引后找真实数据
	hits := searchResults.Hits
	for _, i := range hits {
		fmt.Println(i.Fields)
	}
	wg.Done()
}
