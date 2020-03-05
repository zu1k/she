package source

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // for mysql
	"github.com/zu1k/she/log"
)

type ku12306 struct {
	db *gorm.DB
}

func init() {
	register("ku12306", newKu12306)
}

func newKu12306(info interface{}) Source {
	link := info.(string)
	db, err := gorm.Open("mysql", link)
	if err != nil {
		log.Errorln("Ku12306 db connect err")
		return nil
	}
	return &ku12306{db: db}
}

// GetName return 12306 name
func (k *ku12306) GetName() string {
	return "Ku12306"
}

// Search return result slice from source 12306
func (k *ku12306) Search(key interface{}) (results []Result) {
	num := key.(int)
	log.Infoln("Search Ku12306, key = %d", num)
	return
}
