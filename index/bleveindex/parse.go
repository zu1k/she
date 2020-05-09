package bleveindex

import (
	"io"

	"github.com/zu1k/she/common/tools"
)

type Entity map[string]interface{}

func Parse(filepath string, infoFilePath string, entityChan chan Entity) {
	columnsInfo := ParseFile(infoFilePath).Columns

	reader, err := tools.OpenCSC(filepath)
	if err != nil {
		panic(err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			switch err {
			case io.EOF:
				close(entityChan)
				return
			default:
				continue
			}
		}
		entity := make(Entity)
		for _, column := range columnsInfo {
			data := column.Parse(record[column.Index])
			entity[column.Name] = data
		}
		entityChan <- entity
	}
}
