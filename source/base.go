package source

import (
	"strings"
	"sync"

	"github.com/zu1k/she/common"
	"github.com/zu1k/she/log"
)

type creator func(info interface{}) Source

var (
	creatorMap = make(map[string]creator)
)

type Source interface {
	GetName() string
	Search(key interface{}, resChan chan common.Result, wg *sync.WaitGroup)
}

func Register(name string, c creator) {
	creatorMap[name] = c
}

// NewSource create an Source object by name and return as an Source interface
func NewSource(name string, info interface{}) Source {
	c, ok := creatorMap[strings.ToLower(name)]
	if ok {
		log.Infoln("Init %s source...", name)
		return c(info)
	}
	log.Errorln("Source type not found: %s", name)
	return nil
}
