package main

import (
	"github.com/zu1k/she/cmd"
	_ "github.com/zu1k/she/cmd"
	_ "github.com/zu1k/she/constant"
	_ "github.com/zu1k/she/persistence"
	_ "github.com/zu1k/she/pkg/source/bleveindex"
	_ "github.com/zu1k/she/pkg/source/plaintext"
	_ "github.com/zu1k/she/pkg/source/qqgroup"
)

func main() {
	cmd.Execute()
}
