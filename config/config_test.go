package config

import "testing"

func TestGoPath(t *testing.T) {
	expectedPath := "/home/ben/go"
	if goPath != expectedPath {
		t.Errorf("Expected %s, got %s", expectedPath, goPath)
	}
}

func TestBasePath(t *testing.T) {
	expectedPath := "/home/ben/go/src/github.com/BenLirio/lirio-tools"
	if BasePath != expectedPath {
		t.Errorf("Expected %s, got %s", expectedPath, BasePath)
	}
}
