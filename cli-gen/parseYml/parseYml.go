package parseYml

import (
	"github.com/BenLirio/lirio-tools/cli-gen/types"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

type CommandYml struct {
	Use      string       `json:"use"`
	Short    string       `json:"short"`
	Long     string       `json:"long"`
	Commands []CommandYml `json:"commands"`
}

type RootYml struct {
	Commands []CommandYml `json:"commands"`
}

type CLIYml struct {
	Root RootYml `json:"root"`
}

func distilYmlCommand(cmdYml CommandYml) types.Command {
	var cmd types.Command
	cmd.Use = cmdYml.Use
	cmd.Short = cmdYml.Short
	cmd.Long = cmdYml.Long
	for _, subCmdYml := range cmdYml.Commands {
		cmd.Commands = append(cmd.Commands, distilYmlCommand(subCmdYml))
	}
	return cmd
}

func ParseCLI(filename string) []types.Command {
	var cli CLIYml
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal([]byte(data), &cli)
	var rootCmds []types.Command
	for _, cmdYml := range cli.Root.Commands {
		cmd := distilYmlCommand(cmdYml)
		rootCmds = append(rootCmds, cmd)
	}
	return rootCmds
}
