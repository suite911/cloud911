package database

import (
	"log"

	"github.com/suite911/cloud911/vars"

	"github.com/suite911/maths911/maths"
	"github.com/suite911/query911/query"
	"github.com/suite911/str911/str"
)

func Register(username, email string, scores [3]float64, minor bool, emwho, emhow, emrel string) string {
	if len(scores) != 3 {
		log.Println("len(scores):", len(scores), ":", scores)
		return "#something-went-wrong"
	}
	sc0 := int(65536.0 * scores[0])
	sc1 := int(65536.0 * scores[1])
	sc2 := int(65536.0 * scores[2])
	net := int(65536.0 * scores[0] * scores[1] * scores[2])
	if sc0 > 0xffff {
		sc0 = 0xffff
	}
	if sc1 > 0xffff {
		sc1 = 0xffff
	}
	if sc2 > 0xffff {
		sc2 = 0xffff
	}
	if net > 0xffff {
		net = 0xffff
	}
	q := query.Query{ DB: DB() }
	q.SQL = `INSERT INTO "RegisteredUsers"("un", "email", "conload", "conchange", "consubmit", ` +
		`"captcha", "minor", "emwho", "emhow", "emrel") VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	q.Exec(username, email, sc0, sc1, sc2, net, minor, emwho, emhow, emrel)
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
