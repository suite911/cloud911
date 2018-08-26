package database

import "github.com/suite911/query911/query"

type QueryUser struct {
	rowID int64 `json:"rowid"`
	id    int64 `json:"id"`
	regd  int64 `json:"regd"`
	verd  int64 `json:"verd"`
}

func User(email, username string) (*QueryUser, error) {
	q := query.Query{DB: DB()}
	q.SQL = `SELECT "_ROWID_", "id", "regd", "verd" FROM "RegisteredUsers" WHERE "email" = ? AND "un" = ?;`
	q.Query(email, username)
	if err := q.ErrorLogNow(); err != nil {
		return nil, err
	}
	if !q.NextOrClose() {
		return nil, q.ErrorLogNow() // probably nil, which is what we want: it means no result
	}
	resp := new(QueryUser)
	q.ScanClose(&resp.rowID, &resp.id, &resp.regd, &resp.verd)
	if err := q.ErrorLogNow(); err != nil {
		return nil, err
	}
	return resp, nil
}
