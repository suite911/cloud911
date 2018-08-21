package handlers

import (
	"fmt"
	"log"

	"github.com/suite911/cloud911/droppriv"
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
	if err := droppriv.Drop(); err != nil {
		log.Fatalln(err)
	}
	path := string(ctx.Path())
	if match, tail := str.CaseTrimPrefix(path, "/api"); match && (len(tail) < 1 || tail[0] == '/') {
		API(ctx, tail)
		return
	}
	if ctx.IsPost() {
		Post(ctx)
		return
	}
	redir := false
	for {
		log.Printf("path: %q", path)
		switch {
		case str.CaseTrimSuffixInPlace(&path, "/"):
			continue
		case str.CaseTrimSuffixInPlace(&path, "/index"):
			fallthrough
		case str.CaseTrimSuffixInPlace(&path, ".html"):
			fallthrough
		case str.CaseTrimSuffixInPlace(&path, ".htm"):
			fallthrough
		case str.CaseTrimSuffixInPlace(&path, ".php"):
			redir = true
			continue
		}
		break
	}
	if redir {
		var uri fasthttp.URI
		ctx.URI().CopyTo(&uri)
		uri.SetPath(path)
		ctx.RedirectBytes(uri.FullURI(), 301)
		return
	}
	if c, ok := pages.CompiledPages[path]; ok && c != nil {
		c.Serve(ctx)
		return
	}
	ctx.SetStatusCode(404)
	if c, ok := pages.CompiledPages["404"]; ok && c != nil {
		c.Serve(ctx)
		return
	}
	ctx.SetContentType("text/plain; charset=utf8")
	fmt.Fprintf(ctx, "Not Found\n\n")

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
	return
}
