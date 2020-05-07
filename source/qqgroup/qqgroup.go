package qqgroup

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/zu1k/she/common"
	"github.com/zu1k/she/source"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql" // for sql server
	"github.com/zu1k/she/log"
)

type qqGroup struct {
	name string
	db   *gorm.DB
}

func init() {
	source.Register(source.QQGroup, newQQGroup)
}

func newQQGroup(name string, info interface{}) source.Source {
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
	return &qqGroup{db: db, name: name}
}

func (q *qqGroup) Name() string {
	return q.name
}

func (q *qqGroup) Type() source.Type {
	return source.QQGroup
}

// Search return result slice from source QQGroup
func (q *qqGroup) Search(key interface{}, resChan chan common.Result, wg *sync.WaitGroup) {
	num := key.(int)
	log.Infoln("Search QQGroup, key = %d", num)
	done := &sync.WaitGroup{}
	done.Add(3)
	go q.searchGroupByGroupNum(num, resChan, done)
	go q.searchMemberByGroupNum(num, resChan, done)
	go q.searchMemberByQQNum(num, resChan, done)
	done.Wait()
	wg.Done()
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

func (q *qqGroup) searchMemberByQQNum(qqNum int, resChan chan common.Result, done *sync.WaitGroup) {
	var memberRes []member
	q.db.Where("QQNum=?", qqNum).Find(&memberRes)
	for _, m := range memberRes {
		result := common.Result{
			Source: q.name + ":MemberByQQNum",
			Score:  1,
			Hit:    strconv.Itoa(qqNum),
			Text:   m.String(),
		}
		resChan <- result
	}
	done.Done()
}

func (q *qqGroup) searchMemberByGroupNum(groupNum int, resChan chan common.Result, done *sync.WaitGroup) {
	var memberRes []member
	q.db.Where("GroupNum=?", groupNum).Find(&memberRes)
	for _, m := range memberRes {
		result := common.Result{
			Source: q.name + ":MemberByGroupNum",
			Score:  1,
			Hit:    strconv.Itoa(groupNum),
			Text:   m.String(),
		}
		resChan <- result
	}
	done.Done()
}

func (q *qqGroup) searchGroupByGroupNum(groupNum int, resChan chan common.Result, done *sync.WaitGroup) {
	var groupRes []group
	q.db.Where("GroupNum=?", groupNum).Find(&groupRes)
	for _, m := range groupRes {
		result := common.Result{
			Source: q.name + ":GroupByGroupNum",
			Score:  1,
			Hit:    strconv.Itoa(groupNum),
			Text:   m.String(),
		}
		resChan <- result
	}
	done.Done()
}
