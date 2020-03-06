package source

import (
	"strings"

	"github.com/zu1k/she/log"
)

type creator func(info interface{}) Source

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
func NewSource(name string, info interface{}) Source {
	c, ok := creatorMap[strings.ToLower(name)]
	if ok {
		log.Infoln("Init %s source...", name)
		return c(info)
	}
	log.Errorln("Source type not found: %s", name)
	return nil
}

type Result struct {
	Score int    `json:"score"`
	Hit   string `json:"hit"`
	Text  string `json:"text"`
}
