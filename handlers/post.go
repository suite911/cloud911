package handlers

import (
	"github.com/suite911/cloud911/register"

	"github.com/valyala/fasthttp"
)

var OverridePost func(*fasthttp.RequestCtx)

func Post(ctx *fasthttp.RequestCtx) {
	switch {
	case OverridePost == nil:
		post(ctx)
	default:
		OverridePost(ctx)
	}
}

func post(ctx *fasthttp.RequestCtx) {
	args := ctx.PostArgs()
	actionBytes := args.Peek("action")
	if !utf.Valid(actionBytes) {
		return
	}
	switch action := string(actionBytes); action {
	case "register":
		_/*attempt*/, _/*err*/ = register.Try(ctx)
	}
}
