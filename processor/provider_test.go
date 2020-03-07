package processor

import (
	"fmt"
	"testing"

	"github.com/zu1k/she/source"
	_ "github.com/zu1k/she/source/plaintext"
)

func TestProviderSearch(t *testing.T) {
	sourceList = append(sourceList, source.NewSource("plaintext", "../ku/12306/account.csv"))
	fmt.Println(len(sourceList))
	results := SearchAllSource("山东大学")
	if len(results) > 0 {
		for _, v := range results {
			fmt.Println(v.Text)
		}
	} else {
		t.Errorf("results 未查到")
	}
}
