package cmd

import (
	"fmt"

	"github.com/zu1k/she/pkg/index/filewatch"

	"github.com/spf13/cobra"
	"github.com/zu1k/she/internal/processor"
	"github.com/zu1k/she/internal/route"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run a web api server",
	Long:  `run a web api server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("she web api will serve at \"%s\", with secret \"%s\"\n\n", bindAddr, secret)
		fmt.Println("Init source list...")
		processor.InitSourceList()
		fmt.Println("Success init source list")
		switch mode {
		case "auto":
			fmt.Println("auto mode")
			go filewatch.DoWatch()
		case "manual":
			fmt.Println("manual mode")
		}
		route.Start(bindAddr, secret)
	},
}

var (
	bindAddr string
	secret   string
	mode     string
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringVarP(&bindAddr, "bind", "b", ":10086", "web api bind address")
	serveCmd.PersistentFlags().StringVarP(&secret, "secret", "s", "", "web api access secret")
	serveCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "manual", "serve mode (manual、auto)")
}
