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
		CREATE TABLE IF NOT EXISTS "RegisteredUsers" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"un" TEXT NOT NULL UNIQUE,
			"email" TEXT NOT NULL,
			"pw" BLOB DEFAULT(NULL),
			"regd" INTEGER NOT NULL ` + defaultNow + `,
			"verd" INTEGER NOT NULL DEFAULT(0),
			"bal" INTEGER NOT NULL DEFAULT(0),
			"captcha" REAL NOT NULL
		);
	`
	q.Exec()
	if !q.OK() {
		return q.LastError()
	}
	q.SQL = `
		CREATE UNIQUE INDEX IF NOT EXISTS "idx_RegisteredUsers_un" ON "RegisteredUsers"("un");
	`
	q.Exec()
	if !q.OK() {
		return q.LastError()
	}
	q.SQL = `
		CREATE INDEX IF NOT EXISTS "idx_RegisteredUsers_email" ON "RegisteredUsers"("email");
	`
	q.Exec()
	if !q.OK() {
		return q.LastError()
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
		if !q.OK() {
			return q.LastError()
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
