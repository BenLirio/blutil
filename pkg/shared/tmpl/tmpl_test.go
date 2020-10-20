package tmpl

import "testing"

func TestDoesCompile(t *testing.T) {
	const templateString string = `This is a test template string`
	mustCompile(templateString)
}

func TestCorrectTemplate(t *testing.T) {
	const templateString string = `With a  = {{.Val}}`
	const expectedReturnString string = `With a  = 3`
	type A struct {
		Val int
	}
	a := A{Val: 3}
	returnString, err := Create(templateString, a)
	if err != nil {
		t.Error(err)
	}
	if returnString != expectedReturnString {
		t.Errorf("Expected %s, got %s", expectedReturnString, returnString)
	}
}
