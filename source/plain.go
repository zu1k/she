package source

import (
	"log"

	"github.com/jinzhu/gorm"
)

type plaintext struct {
	db *gorm.DB
}

func init() {
	register("plaintext", newPlain)
}

func (p *plaintext) GetName() string {
	return "Plain"
}

func (p *plaintext) Search(key interface{}) (result []Result) {
	return nil
}

func newPlain() Source {
	db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
	if err != nil {
		log.Println("")
	}
	return &plaintext{db: db}
}
