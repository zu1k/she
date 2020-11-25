package bleveindex

import "strconv"

type ColumnsInfo []Column

// 配置文件中每一项定义需要的字段是哪一列，什么数据类型、列名
type Column struct {
	Index int    `yaml:"index"`
	Type  string `yaml:"type"`
	Name  string `yaml:"name"`
}

func (c Column) Parse(str string) interface{} {
	switch c.Type {
	case "string":
		return str
	case "int":
		intdata, err := strconv.Atoi(str)
		if err != nil {
			return ""
		}
		return intdata
	default:
		return str
	}
}
