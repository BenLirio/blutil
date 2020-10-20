package cmd

type templateStruct struct {
	Use string
}

var templateString string = `
package main

import (
	"github.com/BenLirio/lirio-tools/pkg/cli/{{.Use}}/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute()
}
`
