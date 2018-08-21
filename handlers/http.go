package handlers

import (
	"log"

	"github.com/suite911/cloud911/droppriv"

	"github.com/suite911/str911/str"

	"github.com/valyala/fasthttp"
)

var OverrideHTTP func(*fasthttp.RequestCtx)

func HTTP(ctx *fasthttp.RequestCtx) {
	switch {
	case OverrideHTTP == nil:
		http(ctx)
	default:
		OverrideHTTP(ctx)
	}
}

func http(ctx *fasthttp.RequestCtx) {
	if err := droppriv.Drop(); err != nil {
		log.Fatalln(err)
	}
	if match, tail := str.CaseTrimPrefix(string(ctx.Path()), "/api"); match && (len(tail) < 1 || tail[0] == '/') {
		API(ctx, tail)
		return
	}
	var uri fasthttp.URI
	ctx.URI().CopyTo(&uri)
	uri.SetScheme("https")
	ctx.RedirectBytes(uri.FullURI(), 301) // 301 recommended by Google
}
