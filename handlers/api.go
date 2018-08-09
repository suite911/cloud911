package handlers

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

var OverrideApi func(*fasthttp.RequestCtx, string)

func Api(ctx *fasthttp.RequestCtx, path string) {
	switch {
	case OverrideApi == nil:
		api(ctx, path)
	default:
		OverrideApi(ctx, path)
	}
}

func api(ctx *fasthttp.RequestCtx, path string) {
	ctx.Error("Not Implemented", 501)
}
