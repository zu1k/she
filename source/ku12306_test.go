package source

import (
	"testing"
)

func TestKu12306Search(t *testing.T) {
	link := ""
	ku := newKu12306(link)
	if ku != nil {
	} else {
		t.Errorf("Ku12306 db connect err")
	}
}
