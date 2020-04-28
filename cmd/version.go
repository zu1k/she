package cmd

import (
	"fmt"

	"github.com/zu1k/she/constant"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show she version",
	Long:  `show she version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(constant.Name, "-", constant.Version, "-", constant.BuildTime, "\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
