package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// analyseCmd represents the analyse command
var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "analyse a file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyse called")
	},
}

func init() {
	rootCmd.AddCommand(analyseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
