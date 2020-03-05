package source

import (
	"fmt"
	"testing"
)

func TestQQGroupSearch(t *testing.T) {
	link := "sqlserver://she:she@192.168.254.145:1433?database=QQGroup"
	qq := newQQGroup(link)
	if qq != nil {
		res := qq.Search(11348929)
		fmt.Printf("QQGroup search 11348929 result: length %d\n", len(res))
		if len(res) != 7 {
			t.Errorf("QQGroup search err")
		}
		for _, r := range res {
			fmt.Println(r.Text)
		}
	} else {
		t.Errorf("QQGroup db connect err")
	}
}
