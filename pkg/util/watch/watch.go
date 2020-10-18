package watch

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
        Use:   "watch",
        Short: "run script on file change",
        Long: "long: run script on file change",
        Run: func(cmd *cobra.Command, args []string) {
                Execute()
        },
}

func Execute() {
	fmt.Println("Hello World")
}
