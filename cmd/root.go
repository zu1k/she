package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zu1k/she/constant"
)

var (
	cfgFile  string
	homePath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "she",
	Short: "A brief description of your application",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	ShowAsciiPic()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initHomePath)

	rootCmd.PersistentFlags().StringVar(&homePath, "path", "", "home dir path")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.she.yaml)")
}

func initHomePath() {
	if homePath != "" {
		constant.SetHomeDir(homePath)
	}
}
