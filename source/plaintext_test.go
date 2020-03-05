package source

import (
	"testing"
)

func TestPlaintextSearch(t *testing.T) {
	searchFileContainsStr("../ku/12306/account.csv", "lvxiao98")
}
