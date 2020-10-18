package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// rootCmd.AddCommand(download.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "util",
	Short: "Personal utilities",
}
