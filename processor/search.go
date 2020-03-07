package processor

import (
	"encoding/json"

	"github.com/zu1k/she/common"
)

// Search2Json 搜索返回json结果
func Search2Json(key string) (resultText string) {
	results := SearchAllSource(key)
	resultList, err := json.Marshal(results)
	if err != nil {
		return ""
	}
	resultText = string(resultList)
	return
}

// Search 搜索返回Result结构
func Search(key string) (resultList []common.Result) {
	resultList = SearchAllSource(key)
	return
}
