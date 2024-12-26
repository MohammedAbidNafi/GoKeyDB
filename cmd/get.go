package cmd

import (
	"GoKeyDB/store"
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [db] [key]",
	Short: "Retrieve the value for a key",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		db := args[0]
		key := args[1]
		value, exists := store.Get(db, key)
		if exists {
			fmt.Printf("Value: %s\n", value)
		} else {
			fmt.Println("Key not found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
