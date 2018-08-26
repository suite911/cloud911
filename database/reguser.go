package database

import (
	"github.com/suite911/cloud911/types"

	"github.com/suite911/query911/query"
)

func Me(id int64) (*QueryMe, error) {
	q := query.Query{DB: DB()}
	q.SQL = `SELECT "_ROWID_", ` +
		`"email", "un", "pw", "regd", "verd", "bal", "minor", "emwho", "emhow", "emrel" ` +
		`FROM "RegisteredUsers" WHERE "id" = ?;`
	q.Query(id)
	if err := q.ErrorLogNow(); err != nil {
		return nil, err
	}
	if !q.NextOrClose() {
		return nil, q.ErrorLogNow() // probably nil, which is what we want: it means no result
	}
	resp := new(QueryMe)
	var pw []byte
	var minor int64
	resp.ID = id
	q.ScanClose(
		&resp.RowID, &resp.Email, &resp.Username, &pw, &resp.Registered, &resp.Verified,
		&resp.Balance, &minor, &resp.EmergencyWho, &resp.EmergencyHow, &resp.EmergencyRel,
	)
	resp.HasPassword = len(pw) > 0
	resp.Minor = minor != 0
	if err := q.ErrorLogNow(); err != nil {
		return nil, err
	}
	return resp, nil
}
