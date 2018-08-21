package database

import (
	"errors"
	"unicode/utf8"

	"github.com/suite911/query911/query"

	"github.com/badoux/checkmail"
)

func Register(email []byte) (string redir, err error) {
	if len(email) < 1 {
		return "?email=missing", errors.New("Empty string")
	}
	if !utf8.Valid(email) {
		return "?email=invalid", errors.New("Malformed UTF-8 string")
	}
	email := string(email)
	if err = checkmail.ValidateFormat(email); err != nil {
		return "?email=invalid", err
	}
	if err = checkmail.ValidateHost(email); err != nil {
		return "?email=invalid", err
	}

	q := query.Query{ DB: DB() }
	q.SQL = `INSERT OR IGNORE INTO "RegisteredUsers"("email") VALUES(?);`
	q.Exec(email)
	if !q.OK() {
		return redir, q.LastError()
	}
	return redir, nil
}
