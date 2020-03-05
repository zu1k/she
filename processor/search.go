package processor

import (
	"encoding/json"

	"github.com/zu1k/she/source"
)

func Search2Json(key string) (resultText string) {
	results := source.Search(key)
	resultList, err := json.Marshal(results)
	if err != nil {
		return ""
	}
	resultText = string(resultList)
	return
}

func Search(key string) (resultList []source.Result) {
	resultList = source.Search(key)
	return
}
