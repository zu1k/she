package source

import (
	"log"
	"strconv"

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

type group struct {
	Id         int    `gorm:"column:Id"`
	GroupNum   int    `gorm:"column:GroupNum"`
	Mast       int    `gorm:"column:Mast"`
	CreateDate string `gorm:"column:CreateDate"`
	Title      string `gorm:"column:Title"`
	Class      string `gorm:"column:Class"`
	Summary    string `gorm:"column:Summary"`
}

func (group) TableName() string {
	return "Group"
}

type member struct {
	Id       int    `gorm:"column:Id"`
	QQNum    int    `gorm:"column:QQNum"`
	Nick     string `gorm:"column:Nick"`
	Age      string `gorm:"column:Age"`
	Gender   int    `gorm:"column:Gender"`
	Auth     int    `gorm:"column:Auth"`
	GroupNum int    `gorm:"column:GroupNum"`
}

func (member) TableName() string {
	return "Member"
}

func (m member) String() string {
	return m.Nick
}

func (q *qqGroup) searchMemberByQQNum(qqNum int) (results []Result) {
	var memberRes []member
	q.db.Where("QQNum=?", qqNum).Find(&memberRes)
	for _, m := range memberRes {
		result := Result{
			Score: 1,
			Hit:   strconv.Itoa(qqNum),
			Text:  m.String(),
		}
		results = append(results, result)
	}
	return
}

func (q *qqGroup) searchMemberByGroupNum(groupNum int) {

}

func (q *qqGroup) searchGroupByGroupNum(groupNum int) {

}
