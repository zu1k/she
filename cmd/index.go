package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zu1k/she/pkg/index/bleveindex"
	"github.com/zu1k/she/pkg/index/fullline"
	"github.com/zu1k/she/pkg/index/jiudian2000w"
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
			case "csv":
				if info := *infoFilePath; info == "" {
					fmt.Println("specific info file path")
				} else {
					bleveindex.ParseAndIndex(filePath, info)
				}
				return
			case "jiudian2000w":
				jiudian2000w.ParseAndIndex(filePath)
			case "line":
				fullline.ParseAndIndex(filePath)
			}
		},
	}
	indexEngineType *string
	infoFilePath    *string
)

func init() {
	rootCmd.AddCommand(indexCmd)

	indexEngineType = indexCmd.Flags().StringP("type", "t", "line", "which index engine to use")
	infoFilePath = indexCmd.Flags().StringP("info", "i", "", "info file to use")
}
