package source

import (
	"fmt"
	"sync"

	"github.com/zu1k/she/common"
	"github.com/zu1k/she/log"
)

type creator func(name string, info interface{}) Source

var (
	creatorMap = make(map[Type]creator)
)

type Source interface {
	Name() string
	Type() Type
	Search(key interface{}, resChan chan common.Result, wg *sync.WaitGroup)
}

func Register(sourceType Type, c creator) {
	creatorMap[sourceType] = c
}

// NewSource create an Source object by name and return as an Source interface
func NewSource(name string, stype Type, info interface{}) Source {
	fmt.Println(name, stype, info)
	c, ok := creatorMap[stype]
	if ok {
		log.Infoln("Init bleve index source:  %s", name)
		return c(name, info)
	}
	log.Errorln("Source type not found: %s", name)
	return nil
}
