package processor

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/zu1k/she/pkg/result"

	"github.com/zu1k/she/persistence"
	"github.com/zu1k/she/pkg/source"
)

var sourceList = make([]source.Source, 0)

func AddSource(name string, p source.Type, src string) {
	sourceList = append(sourceList, source.NewSource(name, p, src))
}

// InitSourceList 初始化搜索资源的列表
func InitSourceList() {
	sourceList = append(sourceList, ReadSourceFromDB()...)
}

func ReadSourceFromDB() []source.Source {
	dbsources, err := persistence.FetchAllSource()
	if err != nil {

	}
	var sourceList = make([]source.Source, 0)
	for _, asource := range dbsources {
		sourceList = append(sourceList, source.NewSource(asource.Name, asource.Type, asource.Src))
	}
	return sourceList
}

// Search search all data source
func SearchAllSource(key string) (results []result.Result) {
	wg := &sync.WaitGroup{}
	resChan := make(chan result.Result, 30)
	fmt.Println("len", len(sourceList))
	for _, s := range sourceList {
		if s == nil {
			continue
		}
		wg.Add(1)
		switch s.Type() {
		case source.QQGroup:
			key = strings.ReplaceAll(key, " ", "")
			num, err := strconv.Atoi(key)
			if err != nil {
				wg.Done()
				continue
			}
			go s.Search(num, resChan, wg)
		case source.PlainText:
			go s.Search(key, resChan, wg)
		case source.BleveIndex:
			go s.Search(key, resChan, wg)
		}
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()
	for result := range resChan {
		results = append(results, result)
	}
	return
}
