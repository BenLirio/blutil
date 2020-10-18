/*
Generated
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/BenLirio/lirio-tools/pkg/util/watch"

)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(watch.Cmd)

}

var rootCmd = &cobra.Command{
	Use:   "util",
	Short: "personal utilities",
	Long: "long personal utilities",
}
