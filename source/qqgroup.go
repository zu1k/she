package source

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type qqGroup struct {
	db *gorm.DB
}

func init() {
	register("qqgroup", newQQGroup)
}

func newQQGroup() Source {
	db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
	if err != nil {
		log.Println("")
	}
	return &qqGroup{db: db}
}

func (q *qqGroup) GetName() string {
	return "QQGroup"
}

func (q *qqGroup) Search(key interface{}) (result []Result) {
	return nil
}
