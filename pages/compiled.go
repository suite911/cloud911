package pages

import (
	"text/template"

	"github.com/suite911/error911/onfail"

	"github.com/valyala/fasthttp"
)

// CompiledPages is a map of all of the pages after they are compiled.
var CompiledPages = make(map[string]*CompiledPage)

// Compile compiles each page from Pages and adds it to CompiledPages.
func Compile(defaultShell *template.Template, onFail ...onfail.OnFail) error {
	for k, v := range Pages {
		c, err := v.Compile(defaultShell, onFail...)
		if err != nil {
			return err
		}
		CompiledPages[k] = c
	}
	return nil
}

// CompiledPage is a type representing a compiled page.
type CompiledPage struct {
	Bytes       []byte
	ContentType string
}

// Serve serves the CompiledPage over the network.
func (c *CompiledPage) Serve(ctx *fasthttp.RequestCtx) {
	if len(c.ContentType) > 0 {
		ctx.SetContentType(c.ContentType)
	}
	ctx.Write(c.Bytes)
}
