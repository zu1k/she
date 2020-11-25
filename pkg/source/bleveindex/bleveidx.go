package bleveindex

import (
	"fmt"
	"sync"

	"github.com/zu1k/she/pkg/result"

	"github.com/zu1k/she/pkg/source"

	"github.com/blevesearch/bleve"
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
		fmt.Println("Fail to open bleveindex index file")
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
func (b *Bleveidx) Search(key interface{}, resChan chan result.Result, wg *sync.WaitGroup) {
	str := key.(string)
	fmt.Println("Search bleve index, key=", key)
	query := bleve.NewMatchQuery(str)
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResults, err := b.index.Search(searchRequest)
	if err != nil {
		return
	}
	hits := searchResults.Hits
	for _, i := range hits {
		resChan <- result.Result{
			Source: "bleve",
			Score:  1,
			Hit:    str,
			Text:   fmt.Sprintln(i.Fields),
		}
	}
	wg.Done()
}
