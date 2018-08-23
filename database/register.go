package database

import (
	"log"

	"github.com/suite911/query911/query"
)

func Register(email string) (redir string, err error) {
	q := query.Query{ DB: DB() }
	q.SQL = `INSERT INTO "RegisteredUsers"("email") VALUES(?);`
	q.Exec(email)
	if !q.OK() {
		log.Printf("INSERT: <%T>: \"%v\"", q.LastError(), q.LastError())
		return "#already-registered", q.LastError()
	}
	return redir, nil
}
