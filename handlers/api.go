package handlers

import (
	"path"

	"github.com/suite911/cloud911/database"
	"github.com/suite911/cloud911/types"

	"github.com/suite911/vault911/vault"

	"github.com/valyala/fasthttp"
)

// APIs holds the encrypted APIs
var APIs = make(map[string]func(*fasthttp.RequestCtx, []byte) []byte)

// Endpoints holds the custom endpoints
var Endpoints = make(map[string]func(*fasthttp.RequestCtx))

// API handles API calls
func API(ctx *fasthttp.RequestCtx, path string) {
	p = path.Clean(p)
	if cb, ok := Endpoints[p]; ok {
		cb(ctx)
		return
	}
	if !ctx.IsPost() {
		ctx.Error("Bad Request", 400)
		return
	}
	args := ctx.PostArgs()
	ridBytes := args.Peek("rowid")
	rid := int64(-1)
	if len(ridBytes) > 0 {
		if !utf8.Valid(ridBytes) {
			ctx.Error("Bad Request", 400)
			return
		}
		i, err := ParseInt(string(ridBytes), 10, 64)
		if err != nil {
			ctx.Error("Bad Request", 400)
			return
		}
		rid = i
	}
	idBytes := args.Peek("id")
	id := int64(-1)
	if len(idBytes) > 0 {
		if !utf8.Valid(idBytes) {
			ctx.Error("Bad Request", 400)
			return
		}
		i, err := ParseInt(string(idBytes), 10, 64)
		if err != nil {
			ctx.Error("Bad Request", 400)
			return
		}
		id = i
	}
	var email, username string
	if i < 1 {
		emailBytes := args.Peek("email")
		if len(emailBytes) < 1 || !utf8.Valid(emailBytes) {
			ctx.Error("Bad Request", 400)
			return
		}
		usernameBytes := args.Peek("username")
		if len(usernameBytes) < 1 || !utf8.Valid(usernameBytes) {
			ctx.Error("Bad Request", 400)
			return
		}
		email, username = string(emailBytes), string(usernameBytes)
	}
	switch p {
	case "/":
		if i < 1 {
			ctx.Error("Bad Request", 400)
			return
		}
		q := query.Query{DB: database.DB()}
		q.SQL = `SELECT "key" FROM "Users" WHERE `
		if id > 0 {
			if rid > 0 {
				q.SQL += `"_ROWID_" = ? && `
			}
			q.SQL += `"id" = ?;`
			if rid > 0 {
				q.Query(rid, id)
			} else {
				q.Query(id)
			}
		} else {
			q.SQL += `"email" = ? AND "username" = ?;`
			q.Query(email, username)
		}
		if err := q.Error; err != nil || !q.NextOrClose() {
			ctx.Error("Unauthorized: user not found", 401)
			return
		}
		var key vault.Key
		q.ScanClose(&key)
		if err := q.Error; err != nil || len(key) != 32 {
			ctx.Error("Internal Server Error: unable to get key for user")
			return
		}
		b, err := vault.Recv(ctx, key)
		var apiCall types.APICall
		if err != nil || json.Unmarshal(b, &apiCall) != nil {
			ctx.Error("Bad Request", 400)
			return
		}
		api, ok := APIs[apiCall.API]
		if !ok {
			ctx.Error("Not Implemented", 501)
			return
		}
		ctx.SetStatusCode(200)
		reply := api(ctx, apiCall.Payload)
		http500, err := vault.Reply(ctx, reply, key)
		if err != nil {
			ctx.Error(http500, 500)
			return
		}
		return
	case "register": // set key
		key := args.Peek("key")
		now := time.Now().UTC().Unix()
		q := query.Query{DB: database.DB()}
		q.SQL = `UPDATE "Users" SET "key" = ?, "pwat" = ? WHERE `
		if id > 0 {
			if rid > 0 {
				q.SQL += `"_ROWID_" = ? && `
			}
			q.SQL += `"id" = ? AND "key" = NULL;`
			if rid > 0 {
				q.Exec(key, now, rid, id)
			} else {
				q.Exec(key, now, id)
			}
		} else {
			q.SQL += `"email" = ? AND "username" = ? AND "key" = NULL;`
			q.Exec(key, now, email, username)
		}
		if err := q.Error; err != nil {
			ctx.Error("Unauthorized: user not found or already has a key", 401)
			return
		}
		ctx.SetStatusCode(200)
		return
	default:
		ctx.Error("Not Implemented", 501)
		return
	}
}
