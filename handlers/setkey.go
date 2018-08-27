package handlers

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/suite911/cloud911/database"
	"github.com/suite911/cloud911/types"

	"github.com/suite911/query911/query"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/sha3"
)

func SetKey(ctx *fasthttp.RequestCtx, setKey []byte) {
	var sk types.SetKey
	if err := json.Unmarshal(setKey, &sk); err != nil {
		ctx.Error("Bad Request: must be valid JSON", 400)
		return
	}
	rid, id, key := sk.RowID, sk.ID, sk.Key
	if len(key) != 32 {
		ctx.Error("Bad Request: length of key must be 32", 400)
		return
	}
	dig := sha3.Sum256(key)
	now := time.Now().UTC().Unix()
	q := query.Query{DB: database.DB()}
	q.SQL = `UPDATE "Users" ` +
		`SET "key" = ?, "ses" = ? ` +
		`WHERE `
	if rid > 0 {
		q.SQL += '"_ROWID_" = ? && '
	}
	q.SQL += '"id" = ? && "key" = NULL;'
	if rid > 0 {
		q.Exec(dig, now, rid, id)
	} else {
		q.Exec(dig, now, id)
	}
	if err := q.Error; err != nil {
		ctx.Error("Forbidden: account does not exist or already has a key", 403)
		return
	}
	s := new(types.Session)
	s.RowID = rid
	s.ID = id
	s.LoggedInAt = now
	s.IP = ctx.RemoteIP()
	b, err := json.Marshal(s)
	if err != nil {
		ctx.Error("Internal Server Error: unable to marshal JSON result", 500)
		return
	}
	if _, err := ctx.Write(b); err != nil {
		ctx.SetStatusCode(500)
		return
	}
	ctx.SetStatusCode(200)
}
