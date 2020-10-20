package cmd

import (
	"github.com/BenLirio/lirio-tools/config"
	"github.com/BenLirio/lirio-tools/pkg/shared/tmpl"
	"io/ioutil"
	"os"
	"path"
)

func generateString(name string) (string, error) {
	res, err := tmpl.Create(templateString, templateStruct{Use: name})
	if err != nil {
		return "", err
	}
	return res, nil
}

func GenerateFile(name string) error {
	str, err := generateString(name)
	if err != nil {
		return err
	}
	data := []byte(str)
	os.Mkdir(path.Join(config.BasePath, "cmd", name), 0775)
	err = ioutil.WriteFile(path.Join(config.BasePath, "cmd", name, name+".go"), data, 0644)
	if err != nil {
		return err
	}
	return nil
}
