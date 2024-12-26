package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gokeydb",
	Short: "A CLI for a Key-Value Store",
	Long:  "GoKeyDB is a simple key-value store implemented in Go with SQLite persistence",
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
