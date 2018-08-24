package vars

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

var Script1 string

type Script struct {
}
var Template *template.Template

func init() {
	text := `
`
	var err error
	if Template, err = template.New("Script").Option("missingkey=zero").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Amy").Parse(text)`))
	}

	var b bytes.Buffer
	if err := Template.Execute(&b, Script{
	}); err != nil {
		panic(err)
	}
	Script1 = string(b.Bytes())
}
