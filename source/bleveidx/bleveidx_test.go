package bleveidx

import (
	"fmt"
	"testing"
)

func TestQQGroupSearch(t *testing.T) {
	b := newBleveIdx("../../ku/test1")
	results := b.Search("zhaoyanxia002")
	if len(results) > 0 {
		for _, i := range results {
			fmt.Println(i.Text)
		}
	} else {
		t.Errorf("not found")
	}
}
