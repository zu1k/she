package source

import (
	"strconv"
	"strings"
)

var sourseList = make([]Source, 0)

func InitSourceList() {
	sourseList = append(sourseList, NewSource("qqgroup", "sqlserver://she:she@192.168.254.145:1433?database=QQGroup"))
	sourseList = append(sourseList, NewSource("plaintext", "./ku/12306/account.csv"))
	sourseList = append(sourseList, NewSource("plaintext", "./ku/12306/relation.csv"))
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
		case "PlainText":
			res = s.Search(key)
		}
		results = append(results, res...)
	}
	return
}
