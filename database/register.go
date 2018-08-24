package database

import (
	"log"

	"github.com/suite911/cloud911/vars"

	"github.com/suite911/query911/query"
	"github.com/suite911/str911/str"
)

func Register(username, email string, scores []float64, minor bool, emwho, emhow, emrel string) string {
	if len(scores) != 3 {
		log.Println("len(scores):", len(scores), ":", scores)
		return "#something-went-wrong"
	}
	q := query.Query{ DB: DB() }
	q.SQL = `INSERT INTO "RegisteredUsers"("un", "email", "conload", "conchange", "consubmit", ` +
		`"minor", "emwho", "emhow", "emrel") VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);`
	q.Exec(username, email, scores[0], scores[1], scores[2], minor, emwho, emhow, emrel)
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
