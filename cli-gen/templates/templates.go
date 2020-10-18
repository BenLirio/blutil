package templates

import (
	"fmt"
	"github.com/BenLirio/lirio-tools/cli-gen/types"
	"io/ioutil"
	"os"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CommandTemplateData struct {
	Use            string
	Short          string
	Long           string
	CommandsImport string
	CommandsInit   string
}

func getCommandImport(parent string, command string) string {
	return fmt.Sprintf("\"github.com/BenLirio/lirio-tools/pkg/%s/%s\"\n", parent, command)
}
func getCommandInit(command string) string {
	return fmt.Sprintf("rootCmd.AddCommand(%s.Cmd)\n", command)
}

func WriteCommand(command types.Command) {
	var dest string
	var fileName string
	var templateSrc string
	var err error

	var commandTemplateData CommandTemplateData
	commandTemplateData.Use = command.Use
	commandTemplateData.Short = command.Short
	commandTemplateData.Long = command.Long
	commandTemplateData.CommandsImport = ""
	commandTemplateData.CommandsInit = ""

	for _, subCommand := range command.Commands {
		commandTemplateData.CommandsImport += getCommandImport(commandTemplateData.Use, subCommand.Use)
		commandTemplateData.CommandsInit += getCommandInit(subCommand.Use)
		dest = fmt.Sprintf("../pkg/%s/%s", commandTemplateData.Use, subCommand.Use)
		fileName = fmt.Sprintf("%s.go", subCommand.Use)
		templateSrc = "templates/subCommand.txt"
		writeTemplateFile(dest, fileName, templateSrc, subCommand)
	}
	// cmd/util/util.go
	dest = fmt.Sprintf("../cmd/%s", command.Use)
	fileName = fmt.Sprintf("%s.go", command.Use)
	templateSrc = "templates/rootCmd.txt"
	err = writeTemplateFile(dest, fileName, templateSrc, commandTemplateData)
	check(err)

	// pkg/util/cmd/cmd.go
	dest = fmt.Sprintf("../pkg/%s/cmd", command.Use)
	fileName = "cmd.go"
	templateSrc = "templates/pkgCmd.txt"
	err = writeTemplateFile(dest, fileName, templateSrc, commandTemplateData)
	check(err)
}

func writeTemplateFile(dest string, fileName string, templateSrc string, templateData interface{}) (err error) {
	templateBytes, err := ioutil.ReadFile(templateSrc)
	if err != nil {
		return err
	}
	tmpl, err := template.New("template").Parse(string(templateBytes))
	if err != nil {
		return err
	}
	os.MkdirAll(dest, 0775)
	f, err := os.Create(fmt.Sprintf("%s/%s", dest, fileName))
	if err != nil {
		return err
	}
	err = tmpl.Execute(f, templateData)
	if err != nil {
		return err
	}
	return nil
}
