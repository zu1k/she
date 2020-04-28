package persistence

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // init for sqlite
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite", "data.db")
	if err != nil {
		panic("persistent db init error")
	}
	db.AutoMigrate(&Source{})
}

// Source
type Source struct {
	gorm.Model
	Name string
	Src  string
}

func NewSource(name, src string) (err error) {
	source := Source{
		Name: name,
		Src:  src,
	}
	if db.NewRecord(source) {
		db.Create(&source)
	} else {
		return errors.New("source already exists")
	}
	return
}

func FetchAllSource() (sources []Source, err error) {
	db.Find(&sources)
	if len(sources) == 0 {
		return nil, errors.New("no source found")
	}
	return sources, nil
}

func GetSourceSByName(name string) (sources []Source, err error) {
	db.Where("name LIKE %?%", name).Find(&sources)
	if len(sources) == 0 {
		return nil, errors.New(fmt.Sprintln("no source found, search by name: ", name))
	}
	return sources, nil
}

func DeleteSourceByName(name string) (err error) {
	return
}
