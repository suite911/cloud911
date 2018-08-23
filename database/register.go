package database

import (
	"errors"
	"unicode/utf8"

	"github.com/suite911/query911/query"

	"github.com/badoux/checkmail"
)

func Register(email string) (redir string, err error) {
	q := query.Query{ DB: DB() }
	q.SQL = `INSERT OR IGNORE INTO "RegisteredUsers"("email") VALUES(?);`
	q.Exec(email)
	if !q.OK() {
		return redir, q.LastError()
	}
	return redir, nil
}
