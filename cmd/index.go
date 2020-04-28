package cmd

import (
	"fmt"

	"github.com/zu1k/she/source/jiudian2000w"

	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var (
	indexCmd = &cobra.Command{
		Use:   "index",
		Short: "create index for plain text",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Now creating index...")
			if len(args) == 0 {
				fmt.Println("specific file path")
				return
			}
			filePath := args[0]
			switch *indexEngineType {
			case "bleve":
				return
			case "jiudian2000w":
				jiudian2000w.ParseAndIndex(filePath)
			}
		},
	}
	indexEngineType *string
)

func init() {
	rootCmd.AddCommand(indexCmd)

	indexEngineType = indexCmd.Flags().StringP("type", "t", "bleve", "which index engine to use")
}
