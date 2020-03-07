package processor

import (
	"strconv"
	"strings"
	"sync"

	"github.com/zu1k/she/source"

	"github.com/zu1k/she/common"
)

var sourceList = make([]source.Source, 0)

// InitSourceList 初始化搜索资源的列表
func InitSourceList() {
	sourceList = append(sourceList, source.NewSource("qqgroup", "sqlserver://she:she@192.168.254.145:1433?database=QQGroup"))
	sourceList = append(sourceList, source.NewSource("plaintext", "./ku/12306/account.csv"))
	sourceList = append(sourceList, source.NewSource("plaintext", "./ku/12306/relation.csv"))
}

// Search search all data source
func SearchAllSource(key string) (results []common.Result) {
	wg := &sync.WaitGroup{}
	resChan := make(chan common.Result, 30)
	for _, s := range sourceList {
		if s == nil {
			continue
		}
		wg.Add(1)
		switch s.GetName() {
		case "QQGroup":
			key = strings.ReplaceAll(key, " ", "")
			num, err := strconv.Atoi(key)
			if err != nil {
				wg.Done()
				continue
			}
			go s.Search(num, resChan, wg)
		case "PlainText":
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
