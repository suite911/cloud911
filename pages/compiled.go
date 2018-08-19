package pages

import (
	"bytes"
	"crypto/sha1"
	"math/rand" // yes, for this use case it is secure enough
	"strconv"
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
	ProofOfWork int
}

// Serve serves the CompiledPage over the network.
func (c *CompiledPage) Serve(ctx *fasthttp.RequestCtx) {
	if len(c.ContentType) > 0 {
		ctx.SetContentType(c.ContentType)
	}
	if proofOfWork := c.ProofOfWork; proofOfWork > 0 {
		actual := rand.Uint32() & 0xffff
		challenge := []byte(strconv.Itoa(actual))
		for j := 0; j < proofOfWork; j++ {
			challenge = sha1.Sum(challenge)[:]
		}
		ctx.Write(bytes.Replace(c.Bytes, []byte("__CHALLENGE__"), challenge, -1))
	} else {
		ctx.Write(c.Bytes)
	}
}
