package main

import (
	"github.com/zu1k/she/hub"
	"github.com/zu1k/she/source"
)

func main() {
	source.InitSourceList()
	hub.Start()
}
