package vars

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

var Script1 string

type Script struct {
}
var ScriptTemplate *template.Template

func init() {
	text := `
`
	var err error
	if ScriptTemplate, err = template.New("Script").Option("missingkey=zero").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Amy").Parse(text)`))
	}

	var b bytes.Buffer
	if err := ScriptTemplate.Execute(&b, Script{
	}); err != nil {
		panic(err)
	}
	Script1 = string(b.Bytes())
}
