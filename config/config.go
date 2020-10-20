package config

import (
	"os"
	"path"
)

var goPath string = os.Getenv("GOPATH")
var BasePath string = path.Join(goPath, "src/github.com/BenLirio/lirio-tools")
