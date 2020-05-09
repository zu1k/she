package cmd

import (
	"fmt"

	"github.com/zu1k/she/index/filewatch"

	"github.com/zu1k/she/hub"

	"github.com/spf13/cobra"
	"github.com/zu1k/she/log"
	"github.com/zu1k/she/processor"
)

// serveCmd represents the serve command
var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "run a web api server",
		Long:  `run a web api server.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("she web api will serve at \"%s\", with secret \"%s\"\n\n", *bindAddr, *secret)
			log.Infoln("Init source list...")
			processor.InitSourceList()
			log.Infoln("Success init source list")
			switch *mode {
			case "auto":
				go filewatch.DoWatch()
				hub.Start(*bindAddr, *secret)
			case "manual":
				hub.Start(*bindAddr, *secret)
			}
		},
	}
	bindAddr *string
	secret   *string
	mode     *string
)

func init() {
	rootCmd.AddCommand(serveCmd)
	bindAddr = serveCmd.Flags().StringP("bind", "b", ":10086", "web api bind address")
	secret = serveCmd.Flags().StringP("secret", "s", "", "web api access secret")
	mode = serveCmd.Flags().StringP("mode", "m", "manual", "serve mode (manual、auto)")
}
