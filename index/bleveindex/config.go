package bleveindex

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Info struct {
	Name    string      `yaml:"name"`
	Columns ColumnsInfo `yaml:"columns"`
}

func ParseFile(filepath string) Info {
	i := Info{}
	if f, err := os.Open(filepath); err != nil {
		panic(err)
	} else {
		yaml.NewDecoder(f).Decode(&i)
		return i
	}
}
