package handlers

import (
	"path"
	"strconv"

	"github.com/suite911/cloud911/database"
	"github.com/suite911/cloud911/types"

	"github.com/suite911/vault911/vault"

	"github.com/valyala/fasthttp"
)

// APIs holds the encrypted APIs
var APIs = map[string]func(*fasthttp.RequestCtx, int64, int64, uint64, []byte) []byte{
	"/user": APIUser,
}

// Endpoints holds the custom endpoints
var Endpoints = map[string]func(*fasthttp.RequestCtx){}

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
	rid := peek(args, "rowid")
	if rid == -400 {
		ctx.Error("Bad Request", 400)
		return
	}
	id := peek(args, "id")
	if id == -400 {
		ctx.Error("Bad Request", 400)
		return
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
	flags := peekUint(args, "flags")
	switch p {
	case "/":
		if i < 1 {
			ctx.Error("Bad Request", 400)
			return
		}
		q := query.Query{DB: database.DB()}
		q.SQL = `SELECT "_ROWID_", "id", "key", "flags" FROM "Users" WHERE `
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
		var dbRowID, dbID int64
		var key vault.Key
		var dbFlags uint64
		q.ScanClose(&dbRowID, &dbID, &key, &dbFlags)
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
		reply := api(ctx, dbRowID, dbID, flags & dbFlags, apiCall.Payload)
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

func APIUser(ctx *fasthttp.RequestCtx, rowID, id int64, flags uint64, payload []byte) []byte {
	var identity types.Identity
	if err := json.Unmarshal(payload, &identity); err != nil {
		ctx.Error("Bad Request: unable to unmarshal JSON inner payload", 400)
		return nil
	}
	if identity.ID == id && identity.RowID < 1 && rowID > 0 {
		identity.RowID = rowID
	}
	q := query.Query{DB: database.DB()}
	q.SQL = `SELECT ` +
		`"_ROWID_", "email", "un", "pw", "regd", "vemd", "vidd", ` +
		`"bal", "conload", "conchange", "consubmit", "captcha", ` +
		`"flags", "emwho", "emhow", "emrel" ` +
		`FROM "Users" WHERE `
	if identity.RowID > 0 {
		q.SQL += `"_ROWID_" = ? && `
	}
	q.SQL += `"id" = ?;`
	if identity.RowID > 0 {
		q.Query(identity.RowID, identity.ID)
	} else {
		q.Query(identity.ID)
	}
	if err := q.Error; err != nil || !q.NextOrClose() {
		ctx.Error("Not Found: requested user not found", 404)
		return nil
	}
	var u types.User
	var pw []byte
	var vEmail, vID int64
	q.ScanClose(
		&u.RowID, &u.Email, &u.Username, &pw, &u.Registered, &vEmail, &vID,
		&u.Balance, &u.Captcha1, &u.Captcha2, &u.Captcha3, &u.Captchas,
		&u.Flags, &u.EmergencyWho, &u.EmergencyHow, &u.EmergencyRel,
	)
	if err := q.LogNow(); err != nil {
		ctx.Error("Internal Server Error: unable to scan returned database row", 500)
		return nil
	}
	u.HasPassword = len(pw) > 0
	u.HasVerifiedEmail = vEmail > 0
	u.HasVerifiedIdentity = vID > 0
	if !flags.Any(types.Staff|types.Admin) {
		u.Captcha1 = -1
		u.Captcha2 = -1
		u.Captcha3 = -1
		u.Captchas = -1
		if identity.ID != id {
			u.Email = ""
			u.Username = ""
			u.Registered = 0
			u.Balance = -1
			u.Flags = 0
			u.EmergencyWho = ""
			u.EmergencyHow = ""
			u.EmergencyRel = ""
		}
	}
	b, err := json.Marshal(u)
	if err != nil {
		ctx.Error("Internal Server Error: unable to marshal JSON result", 500)
		return nil
	}
	return b
}

func peek(args *fasthttp.Args, k string) int64 {
	bytes := args.Peek(k)
	if len(bytes) > 0 {
		if !utf8.Valid(bytes) {
			return -400
		}
		i, err := ParseInt(string(bytes), 10, 64)
		if err != nil {
			return -400
		}
		return i
	}
	return -1
}

func peekUint(args *fasthttp.Args, k string) uint64 {
	bytes := args.Peek(k)
	if len(bytes) > 0 {
		if !utf8.Valid(bytes) {
			return 0
		}
		i, err := ParseUint(string(bytes), 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}
