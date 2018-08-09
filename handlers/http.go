package handlers

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

var OverrideHttp func(*fasthttp.RequestCtx)

func Http(ctx *fasthttp.RequestCtx) {
	switch {
	case OverrideHttp == nil:
		root(ctx)
	default:
		OverrideHttp(ctx)
	}
}

func http(ctx *fasthttp.RequestCtx) {
	path := strings.Split(ctx.Path, "/")
	if len(path) > 0 {
		switch simp := str.Simp(path[0]); simp {
		case "api":
			Api(ctx)
			return
		}
	}
	var uri fasthttp.URI
	ctx.URI.CopyTo(&uri)
	uri.SetScheme("https")
	ctx.RedirectBytes(uri.FullURI(), 301) // 301 recommended by Google
}
