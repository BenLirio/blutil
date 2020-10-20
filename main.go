package main

import (
	"github.com/BenLirio/lirio-tools/cmd"

	"sigs.k8s.io/yaml"
)

type Command struct {
	Use   string `json:"use"`
	Short string `json:"short"`
}

type Root struct {
	Commands []Command `json:"commands"`
}

func GenerateCommands(cliYaml string) (err error) {
	var cmdRoot Root
	err = yaml.Unmarshal([]byte(cliYaml), &cmdRoot)
	if err != nil {
		panic(err)
	}
	for _, command := range cmdRoot.Commands {
		cmd.GenerateFile(command.Use)
	}
	return nil

}

func main() {
	GenerateCommands(`
commands:
  - use: "util"
    short: "personal utils"
`)
}
