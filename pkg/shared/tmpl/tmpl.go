package tmpl

import (
	"bytes"
	"text/template"
)

func mustCompile(ts string) (t *template.Template) {
	t, err := template.New("test").Parse(ts)
	if err != nil {
		panic(err)
	}
	return t
}

func Create(ts string, data interface{}) (string, error) {
	t := mustCompile(ts)
	var b bytes.Buffer
	err := t.Execute(&b, data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
