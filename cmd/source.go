package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "manage sources",
	Long:  `manage all the she sources.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("source called")
	},
}

var sourceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add source",
	Long:  `add source.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("？")
		} else {
			fmt.Println("source add called", args[0])
		}
	},
}

var sourceDelCmd = &cobra.Command{
	Use:   "del",
	Short: "delete source",
	Long:  `delete source.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("source delete called")
	},
}

var (
	all *bool
)

func init() {
	rootCmd.AddCommand(sourceCmd)
	sourceCmd.AddCommand(sourceAddCmd)
	sourceCmd.AddCommand(sourceDelCmd)
	all = sourceDelCmd.Flags().BoolP("all", "a", false, "delete ALL the sources")
}
