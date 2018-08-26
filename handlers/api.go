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
		switch path {
		case "/login":
		case "/user":
		}
	} else {
		switch path {
		}
	}
	ctx.Error("Not Implemented", 501)
}
