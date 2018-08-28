package database

import (
	"database/sql"
	"sync"

	"github.com/suite911/cloud911/vars"

	"github.com/suite911/query911/query"

	_ "github.com/mattn/go-sqlite3"
)

func Create() error {
	if err := Open(vars.Pass.DataBase); err != nil {
		return err
	}

	q := query.Query{DB: DB()}
	q.SQL = `
		CREATE TABLE IF NOT EXISTS "Users" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"email" TEXT NOT NULL,
			"un" TEXT NOT NULL,
			"key" BLOB DEFAULT(NULL),
			"regd" INTEGER NOT NULL ` + defaultNow + `,
			"vemd" INTEGER NOT NULL DEFAULT(0),
			"pwat" INTEGER NOT NULL DEFAULT(0),
			"vidd" INTEGER NOT NULL DEFAULT(0),
			"bal" INTEGER NOT NULL DEFAULT(0),
			"conload" INTEGER NOT NULL,
			"conchange" INTEGER NOT NULL,
			"consubmit" INTEGER NOT NULL,
			"captcha" INTEGER NOT NULL,
			"flags" INTEGER NOT NULL,
			-- Emergency Contact --
			"emwho" TEXT NOT NULL,
			"emhow" TEXT NOT NULL,
			"emrel" TEXT NOT NULL,
			UNIQUE("email", "un")
		);
	`
	q.Exec()
	if err := q.ErrorLogNow(); err != nil {
		return err
	}
	q.SQL = `CREATE UNIQUE INDEX IF NOT EXISTS "idx_Users_email_un" ON ` +
		`"Users"("email", "un");`
	q.Exec()
	if err := q.ErrorLogNow(); err != nil {
		return err
	}
	q.SQL = `CREATE INDEX IF NOT EXISTS "idx_Users_email" ON "Users"("email");`
	q.Exec()
	if err := q.ErrorLogNow(); err != nil {
		return err
	}
	if vars.Pass.FeatureUserProfiles {
		q.SQL = `
			CREATE TABLE IF NOT EXISTS "UserProfiles" (
				"id" INTEGER NOT NULL PRIMARY KEY,
				"lname" TEXT NOT NULL DEFAULT(''),
				"lkana" TEXT NOT NULL DEFAULT(''),
				"gname" TEXT NOT NULL DEFAULT(''),
				"gkana" TEXT NOT NULL DEFAULT(''),
				"names" TEXT NOT NULL DEFAULT(''),
				"gender" TEXT NOT NULL DEFAULT('')
				-- much more to come, but I don't want to work on this right now
			);
		`
		q.Exec()
		if err := q.ErrorLogNow(); err != nil {
			return err
		}
	}
	return nil
}

func DB() *sql.DB {
	return db
}

func Open(path string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if db == nil {
		var err error
		if db, err = sql.Open("sqlite3", path); err != nil {
			return err
		}
	}
	return nil
}

const defaultNow = `DEFAULT(CAST(strftime('%s', 'now') AS INTEGER))`

var db *sql.DB
var mutex sync.Mutex
