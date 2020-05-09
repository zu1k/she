package persistence

import (
	"errors"
	"fmt"

	"github.com/zu1k/she/source"

	"github.com/jinzhu/gorm"
)

// Source
type Source struct {
	gorm.Model
	Name string
	Type source.Type
	Src  string
}

func NewSource(name string, sourceType source.Type, src string) (err error) {
	source := Source{
		Name: name,
		Type: sourceType,
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

func DeleteAllSource() {
	db.Delete(&Source{})
}
