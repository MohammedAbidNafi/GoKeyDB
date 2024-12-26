package cmd

import (
	"GoKeyDB/store"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [db] [key]",
	Short: "Delete a key-value pair",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		db := args[0]
		key := args[1]
		if err := store.Delete(db, key); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Key deleted successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
