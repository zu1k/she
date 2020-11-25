package cmd

import (
	"fmt"

	"github.com/zu1k/she/persistence"

	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "manage sources",
	Long:  `manage all the she sources.`,
	Run: func(cmd *cobra.Command, args []string) {
		sources, err := persistence.FetchAllSource()
		if err != nil {

		}
		fmt.Printf("Name\t\tType\t\tSource\n")
		for _, source := range sources {
			fmt.Printf("%s\t%s\t%s\n", source.Name, source.Type.String(), source.Src)
		}
		fmt.Printf("\nTotal: %d\n", len(sources))
	},
}

var sourceListCmd = &cobra.Command{
	Use:   "list",
	Short: "list source",
	Long:  `list source.`,
	Run: func(cmd *cobra.Command, args []string) {
		sources, err := persistence.FetchAllSource()
		if err != nil {

		}
		fmt.Printf("Name\t\tType\t\tSource\n")
		for _, source := range sources {
			fmt.Printf("%s\t%s\t%s\n", source.Name, source.Type.String(), source.Src)
		}
		fmt.Printf("\nTotal: %d\n", len(sources))
	},
}

var sourceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add source",
	Long:  `add source.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("暂未实现")
	},
}

var sourceDelCmd = &cobra.Command{
	Use:   "del",
	Short: "delete source",
	Long:  `delete source.`,
	Run: func(cmd *cobra.Command, args []string) {
		if *all {
			persistence.DeleteAllSource()
			fmt.Println("all the sources has been deleted")
		} else {
			fmt.Println("NONE the sources has been deleted")
		}
	},
}

var (
	all *bool
)

func init() {
	rootCmd.AddCommand(sourceCmd)
	sourceCmd.AddCommand(sourceAddCmd)
	sourceCmd.AddCommand(sourceDelCmd)
	sourceCmd.AddCommand(sourceListCmd)
	all = sourceDelCmd.Flags().BoolP("all", "a", false, "manage ALL the sources")
}
