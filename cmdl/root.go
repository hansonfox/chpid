package cmdl

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(randCmd)
	rootCmd.AddCommand(validCmd)
	rootCmd.AddCommand(parseCmd)
}
