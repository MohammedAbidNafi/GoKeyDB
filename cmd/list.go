package cmd

import (
	"GoKeyDB/store"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [db]",
	Short: "List all key-value pairs",
	Run: func(cmd *cobra.Command, args []string) {
		db := args[0]
		fmt.Println("Key-Value Pairs:")
		store.List(db)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
