package bleveindex

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
	name  string
}

func init() {
	source.Register(source.BleveIndex, OpenBleveIdx)
}

func OpenBleveIdx(name string, info interface{}) source.Source {
	path := info.(string)
	index, err := bleve.Open(path)
	if err != nil {
		log.Errorln("Fail to open bleveindex index file")
		return nil
	}
	return &Bleveidx{index: index, name: name}
}

func (b *Bleveidx) Name() string {
	return b.name
}

func (b *Bleveidx) Type() source.Type {
	return source.BleveIndex
}

// Search return result slice from source Bleveidx
func (b *Bleveidx) Search(key interface{}, resChan chan common.Result, wg *sync.WaitGroup) {
	str := key.(string)
	log.Infoln("Search bleve index, key = %s", key)
	query := bleve.NewMatchQuery(str)
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResults, err := b.index.Search(searchRequest)
	if err != nil {
		return
	}
	hits := searchResults.Hits
	for _, i := range hits {
		result := common.Result{
			Source: "bleve",
			Score:  1,
			Hit:    str,
			Text:   fmt.Sprintln(i.Fields),
		}
		resChan <- result
	}
	wg.Done()
}
