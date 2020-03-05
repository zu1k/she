package source

import (
	"strconv"
	"strings"
)

var sourseList = make([]Source, 0)

func InitSourceList() {
	sourseList = append(sourseList, NewSource("qqgroup"))
}

// Search search all data source
func Search(key string) (results []Result) {
	var res []Result
	for _, s := range sourseList {
		if s == nil {
			continue
		}
		switch s.GetName() {
		case "QQGroup":
			key = strings.ReplaceAll(key, " ", "")
			num, err := strconv.Atoi(key)
			if err != nil {
				continue
			}
			res = s.Search(num)
		}
		results = append(results, res...)
	}
	return
}
