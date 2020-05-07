package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // init for sqlite
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("persistent db init error")
	}
	db.AutoMigrate(&Source{})
}
