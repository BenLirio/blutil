package main

import (
	"github.com/BenLirio/lirio-tools/cli-gen/parseYml"
	"github.com/BenLirio/lirio-tools/cli-gen/templates"
)

func main() {
	rootCmds := parseYml.ParseCLI("test.yml")
	for _, command := range rootCmds {
		templates.WriteCommand(command)
	}
}
