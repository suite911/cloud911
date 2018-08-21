package handlers

import (
	"fmt"
	"log"

	"github.com/suite911/cloud911/droppriv"
	"github.com/suite911/cloud911/pages"

	"github.com/suite911/str911/str"

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
	switch action := args.Peek("action"); action {
	case "register":
		_/*attempt*/, _/*err*/ := register.Try(ctx)
	}
}
