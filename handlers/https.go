package handlers

import (
	"fmt"
	"io"
	"log"
	"path"
	"strings"

	"github.com/suite911/cloud911/droppriv"
	"github.com/suite911/cloud911/pages"
	"github.com/suite911/cloud911/vars"

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
	p := string(ctx.Path())
	if match, tail := str.CaseTrimPrefix(p, "/api"); match && (len(tail) < 1 || tail[0] == '/') {
		API(ctx, tail)
		return
	}
	if ctx.IsPost() {
		Post(ctx)
		return
	}
	p = path.Clean(p)
	redir := false
	for {
		switch {
		case str.CaseTrimSuffixInPlace(&p, "/index"):
			fallthrough
		case str.CaseTrimSuffixInPlace(&p, ".html"):
			fallthrough
		case str.CaseTrimSuffixInPlace(&p, ".htm"):
			fallthrough
		case str.CaseTrimSuffixInPlace(&p, ".php"):
			redir = true
			continue
		}
		break
	}
	if len(p) < 1 {
		p = "/"
	}
	if redir {
		var uri fasthttp.URI
		ctx.URI().CopyTo(&uri)
		uri.SetPath(p)
		ctx.RedirectBytes(uri.FullURI(), 301)
		return
	}

	// Look in important (not overridable) pages first.

	switch strings.ToLower(p) {
	case "/0.js":
		ctx.SetContentType("application/javascript; charset=utf8")
		io.WriteString(ctx, "var amy_reCAPTCHAv3SiteKey = \""+vars.CaptchaSiteKey+"\";")
		return
	}

	// Not found in important pages.  Look in user's custom pages pages next.

	if c, ok := pages.CompiledPages[p]; ok && c != nil {
		c.Serve(ctx)
		return
	}

	// Not found in user's custom pages.  Look in predefined pages next.

	switch strings.ToLower(p) {
	case "/1.css":
		ctx.SetContentType("text/css; charset=utf8")
		io.WriteString(ctx, vars.Style1)
		return
	case "/1.js":
		ctx.SetContentType("application/javascript; charset=utf8")
		io.WriteString(ctx, vars.Script1)
		return
	}

	// Not found in predefined pages either.  Log the 404.

	log.Printf("404 of %q by %q on %q", p, ctx.RemoteIP(), ctx.UserAgent())

	// Now look for custom 404 page.

	ctx.SetStatusCode(404)
	if c, ok := pages.CompiledPages["404"]; ok && c != nil {
		c.Serve(ctx)
		return
	}

	// No custom 404 page.  Send the ugly default.

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
