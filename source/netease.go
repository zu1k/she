package source

import (
	"log"

	"github.com/jinzhu/gorm"
)

type netease struct {
	db *gorm.DB
}

func init() {
	register("netease", newNetease)
}

func (n *netease) GetName() string {
	return "QQGroup"
}

func (n *netease) Search(key interface{}) (result []Result) {
	return nil
}

func newNetease() Source {
	db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
	if err != nil {
		log.Println("")
	}
	return &netease{db: db}
}
