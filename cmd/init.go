package cmd

import (
	"GoKeyDB/store"
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [dbname]",
	Short: "Initialize the DB",
	Run: func(cmd *cobra.Command, args []string) {
		dbname := args[0]
		fmt.Println("Initialise the DB")
		err := store.Initialize(dbname)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("DB initialized successfully.")

		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
