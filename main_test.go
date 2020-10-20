package main

import "testing"
import "os"

func TestGenerateCommands(t *testing.T) {
	var yamlData string = `
commands:
  - use: "sample"
    short: "personal utils"
`
	err := GenerateCommands(yamlData)
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat("cmd/sample/sample.go"); os.IsNotExist(err) {
		t.Error("cmd/sample/sample.go does not exist")
	}
	os.RemoveAll("cmd/sample")
}
