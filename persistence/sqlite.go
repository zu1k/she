package persistence

import (
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // init for sqlite
	"github.com/zu1k/she/constant"
)

var db *gorm.DB

func init() {
	var err error
	dbpath := filepath.Join(constant.Path.HomeDir(), "data.db")
	db, err = gorm.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Source{})
}
