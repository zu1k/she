package processor

import (
	"encoding/json"

	"github.com/zu1k/she/source"
)

// Search2Json 搜索返回json结果
func Search2Json(key string) (resultText string) {
	results := source.Search(key)
	resultList, err := json.Marshal(results)
	if err != nil {
		return ""
	}
	resultText = string(resultList)
	return
}

// Search 搜索返回Result结构
func Search(key string) (resultList []source.Result) {
	resultList = source.Search(key)
	return
}
