package cmd

import (
	"fmt"
	"os"

	"github.com/zu1k/she/hub"
	"github.com/zu1k/she/log"
	"github.com/zu1k/she/source"
)

const help = `usage: she <command> [<args>]
The most commonly used she commands are:
  run     Start she server
Run 'she <command> -h' for more information on a command.`

func printHelpAndExit() {
	fmt.Println(help)
	os.Exit(0)
}

func She() {
	args := os.Args
	if len(os.Args) <= 1 {
		printHelpAndExit()
	}
	subCommand := os.Args[1]
	os.Args = args[1:]
	switch subCommand {
	case "run":
		log.Infoln("Init source list...")
		source.InitSourceList()
		log.Infoln("Success init source list")
		hub.Start()
	default:
		log.Infoln("do Nothing")
		return
	}
}
