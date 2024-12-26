package cmd

import (
	"GoKeyDB/store"
	"fmt"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put [db] [key] [value]",
	Short: "Add or update a key-value pair",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		db := args[0]
		key := args[1]
		value := args[2]
		if err := store.Put(db, key, value); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Key added/updated successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
