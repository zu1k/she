package source

import "strings"

type creator func() Source

var (
	creatorMap = make(map[string]creator)
)

type Source interface {
	GetName() string
	Search(key interface{}) (result []Result)
}

func register(name string, c creator) {
	creatorMap[name] = c
}

// NewSource create an Source object by name and return as an Source interface
func NewSource(name string) Source {
	c, ok := creatorMap[strings.ToLower(name)]
	if ok {
		return c()
	}
	return nil
}

type Result struct {
	Score int
	Hit   string
	Text  string
}
