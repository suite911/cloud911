package handlers

import (
	"fmt"

	"github.com/suite911/cloud911/pages"

	"github.com/suite911/str911/str"

	"github.com/valyala/fasthttp"
)

var OverrideHTTPS func(*fasthttp.RequestCtx)

func HTTPS(ctx *fasthttp.RequestCtx) {
	switch {
	case OverrideHTTPS == nil:
		https(ctx)
	default:
		OverrideHTTPS(ctx)
	}
}

func https(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	if match, tail := str.CaseHasPrefix(path, "/api"); match {
		API(ctx, tail)
		return
	}
	b, ok := PageBytes[path]
	if !ok {
		fmt.Fprintf(ctx, "Hello, world!\n\n")

		fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
		fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
		fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
		fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
		fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
		fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
		fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
		fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
		fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
		fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

		fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

		ctx.SetContentType("text/plain; charset=utf8")

		ctx.Error("Not Found", 404)
		return
	}

	ctx.Write(b)
	return
}
