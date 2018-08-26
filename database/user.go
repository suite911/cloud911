package database

import (
	"log"

	"github.com/suite911/cloud911/vars"

	"github.com/suite911/query911/query"
	"github.com/suite911/str911/str"
)

type QueryUser struct {
	rowID int64 `json:"rowid"`
	id    int64 `json:"id"`
	regd  int64 `json:"regd"`
	verd  int64 `json:"verd"`
	minor bool  `json:"minor"`
}

func User(email, username string) (*QueryUser, error) {
	q := query.Query{DB: DB()}
	q.SQL = `SELECT "_ROWID_", "id", "regd", "verd", "minor" FROM "RegisteredUsers" WHERE "email" = ? AND "un" = ?;`
	q.Query(email, username)
	if !q.OK() {
		return nil, q.LastError()
	}
	if !q.NextOrClose() {
		return nil, q.LastError() // probably nil, which is what we want: it means no result
	}
	resp := new(QueryUser)
	var minor int64
	q.ScanClose(&resp.rowID, &resp.id, &resp.regd, &resp.verd, &minor)
	resp.minor = minor != 0
	if !q.OK() {
		return nil, q.LastError()
	}
	return resp, nil
}
