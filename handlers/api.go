package handlers

import (
	"github.com/valyala/fasthttp"
)

var OverrideAPI func(*fasthttp.RequestCtx, string)
var APIOverrides = make(map[string]func(*fasthttp.RequestCtx))

func API(ctx *fasthttp.RequestCtx, path string) {
	switch {
	case OverrideAPI == nil:
		api(ctx, path)
	default:
		OverrideAPI(ctx, path)
	}
}

type APILogInRequest struct {
}

type APIUserRequest struct {
}

func api(ctx *fasthttp.RequestCtx, path string) {
	if cb, ok := APIOverrides[path]; ok {
		cb(ctx)
		return
	}
	if ctx.IsPost() {
		args := ctx.PostArgs()
		var resp fasthttp.Args
		switch path {
		case "/login":
			rowid := args.Peek("rowid")
			id := args.Peek("id")
			key := args.Peek("key")
			rand := args.Peek("rand")
			dig := args.Peek("dig")
			_ = rowid
			_ = id
			_ = key
			_ = rand
			_ = dig
		case "/user":
			email := args.Peek("email")
			username := args.Peek("username")
			_ = email
			_ = username
			// resp.Set()
		}
	} else {
		switch path {
		}
	}
	ctx.Error("Not Implemented", 501)
}
