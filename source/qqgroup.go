package source

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql" // for sql server
	"github.com/zu1k/she/log"
)

type qqGroup struct {
	db *gorm.DB
}

func init() {
	register("qqgroup", newQQGroup)
}

func newQQGroup(info interface{}) Source {
	link := info.(string)
	if !strings.Contains(link, "dial+timeout=") {
		if strings.Contains(link, "?") {
			link += "&dial+timeout=5"
		} else {
			link += "?dial+timeout=5"
		}
	}
	fmt.Println(link)
	db, err := gorm.Open("mssql", link)
	if err != nil {
		log.Errorln("QQGroup db connect err")
		return nil
	}
	return &qqGroup{db: db}
}

// GetName return qqgroup name
func (q *qqGroup) GetName() string {
	return "QQGroup"
}

// Search return result slice from source QQGroup
func (q *qqGroup) Search(key interface{}) (results []Result) {
	num := key.(int)
	log.Infoln("Search QQGroup, key = %d", num)
	results = append(results, q.searchMemberByQQNum(num)...)
	results = append(results, q.searchMemberByGroupNum(num)...)
	results = append(results, q.searchGroupByGroupNum(num)...)
	return
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

func (g group) String() string {
	return fmt.Sprintf("Group\t Id: %d, GroupNum: %d, Mast: %d, Title: %s, Class: %s, CreateDate: %s, Summary: %s",
		g.Id, g.GroupNum, g.Mast, g.Title, g.Class, g.CreateDate, g.Summary)
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
	return fmt.Sprintf("Member\t Id: %d, QQNum: %d, Nick: %s, Age: %s, Gender: %d, Auth: %d, GroupNum: %d",
		m.Id, m.QQNum, m.Nick, m.Age, m.Gender, m.Auth, m.GroupNum)
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

func (q *qqGroup) searchMemberByGroupNum(groupNum int) (results []Result) {
	var memberRes []member
	q.db.Where("GroupNum=?", groupNum).Find(&memberRes)
	for _, m := range memberRes {
		result := Result{
			Score: 1,
			Hit:   strconv.Itoa(groupNum),
			Text:  m.String(),
		}
		results = append(results, result)
	}
	return
}

func (q *qqGroup) searchGroupByGroupNum(groupNum int) (results []Result) {
	var groupRes []group
	q.db.Where("GroupNum=?", groupNum).Find(&groupRes)
	for _, m := range groupRes {
		result := Result{
			Score: 1,
			Hit:   strconv.Itoa(groupNum),
			Text:  m.String(),
		}
		results = append(results, result)
	}
	return
}
