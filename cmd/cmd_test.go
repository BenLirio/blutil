package cmd

import "os"
import "testing"

func TestCorrectTemplate(t *testing.T) {
	var expectedString string = `
package main

import (
	"github.com/BenLirio/lirio-tools/pkg/sample/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute()
}
`
	result, err := generateString("sample")
	if err != nil {
		t.Error(err)
	} else if result != expectedString {
		t.Errorf("Expected '%s', got '%s'", expectedString, result)
	}
}

func TestFileCreations(t *testing.T) {
	err := GenerateFile("sample")
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll("sample")
}
