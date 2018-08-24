package database

import (
	"log"

	"github.com/suite911/cloud911/vars"

	"github.com/suite911/query911/query"
	"github.com/suite911/str911/str"
)

func Register(username, email string, minor bool, emwho, emhow, emrel string, captcha float32) string {
	q := query.Query{ DB: DB() }
	q.SQL = `INSERT INTO "RegisteredUsers"("username", "email", "captcha") VALUES(?, ?);`
	q.Exec(username, email, captcha)
	if !q.OK() {
		err := q.LastError()
		if str.CaseHasPrefix(err.Error(), "unique") {
			if url := vars.Pass.AlreadyRegistered; len(url) > 0 {
				return url
			}
			return "#already-registered"
		}
		log.Println(err)
		return "#database-error"
	}
	if url := vars.Pass.Registered; len(url) > 0 {
		return url
	}
	return "#registered"
}
