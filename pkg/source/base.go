package source

import (
	"fmt"
	"sync"

	"github.com/zu1k/she/pkg/result"
)

type creator func(name string, info interface{}) Source

var creatorMap = make(map[Type]creator)

type Source interface {
	Name() string
	Type() Type
	Search(key interface{}, resChan chan result.Result, wg *sync.WaitGroup)
}

func Register(sourceType Type, c creator) {
	creatorMap[sourceType] = c
}

// NewSource create an Source object by name and return as an Source interface
func NewSource(name string, stype Type, info interface{}) Source {
	fmt.Println(name, stype, info)
	c, ok := creatorMap[stype]
	if ok {
		fmt.Println("Init bleve index source:", name)
		return c(name, info)
	} else {
		fmt.Println("Source type not found:", name)
		return nil
	}
}
