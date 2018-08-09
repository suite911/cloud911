package handlers

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

var OverrideApi func(*fasthttp.RequestCtx)

func Api(ctx *fasthttp.RequestCtx) {
	switch {
	case OverrideApi == nil:
		api(ctx)
	default:
		OverrideApi(ctx)
	}
}

func api(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, from the API!")
}
