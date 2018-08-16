package handlers

import (
	"github.com/valyala/fasthttp"
)

var OverrideAPI func(*fasthttp.RequestCtx, string)

func API(ctx *fasthttp.RequestCtx, path string) {
	switch {
	case OverrideAPI == nil:
		api(ctx, path)
	default:
		OverrideAPI(ctx, path)
	}
}

func api(ctx *fasthttp.RequestCtx, path string) {
	ctx.Error("Not Implemented", 501)
}
