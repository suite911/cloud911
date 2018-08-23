package database

import (
	"log"

	"github.com/suite911/cloud911/vars"

	"github.com/suite911/query911/query"
)

func Register(email string) string {
	q := query.Query{ DB: DB() }
	q.SQL = `INSERT INTO "RegisteredUsers"("email") VALUES(?);`
	q.Exec(email)
	if !q.OK() {
		err := q.LastError()
		if str.CaseHasPrefix(err.Error(), "unique") {
			if url := vars.AlreadyRegistered; len(url) > 0 {
				return url
			}
			return "#already-registered"
		}
		log.Println(err)
		return "#database-error"
	}
	if url := vars.Registered; len(url) > 0 {
		return url
	}
	return "#registered"
}
