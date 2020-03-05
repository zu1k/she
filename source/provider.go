package source

var sourseList = make([]Source, 0)

func init() {
	sourseList = append(sourseList, NewSource("qqgroup"))
}

// Search search all data source
func Search(key string) (results []Result) {
	for _, s := range sourseList {
		res := s.Search(key)
		results = append(results, res...)
	}
	return
}
