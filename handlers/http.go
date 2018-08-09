package handlers

import (
	"github.com/amy911/amy911/str"

	"github.com/valyala/fasthttp"
)

var OverrideHttp func(*fasthttp.RequestCtx)

func Http(ctx *fasthttp.RequestCtx) {
	switch {
	case OverrideHttp == nil:
		http(ctx)
	default:
		OverrideHttp(ctx)
	}
}

func http(ctx *fasthttp.RequestCtx) {
	if str.CaseHasPrefix(ctx.Path(), "/api") {
		Api(ctx)
	}
	var uri fasthttp.URI
	ctx.URI.CopyTo(&uri)
	uri.SetScheme("https")
	ctx.RedirectBytes(uri.FullURI(), 301) // 301 recommended by Google
}
