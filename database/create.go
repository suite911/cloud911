package database

import (
	"database/sql"

	"github.com/suite911/query911/query"
)

const defaultNow = `DEFAULT(CAST(strftime('%s', 'now') AS INTEGER))`

DB *sql.DB

func Create() error {
	if DB == nil {
		var err error
		DB, err = sql.Open("sqlite3", vars.DataBase)
		if err != nil {
			return err
		}
	}

	q := query.Query{DB: DB}
	q.SQL = `
		CREATE TABLE IF NOT EXISTS "RegisteredUsers" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"email" TEXT NOT NULL UNIQUE,
			"pw" BLOB, -- specifically can be NULL
			"credits" INTEGER NOT NULL DEFAULT(0),
			"ts" INTEGER NOT NULL ` + defaultNow + `
		);
	`
	q.Exec()
	if !q.Ok() {
		return q.LastError()
	}
	q.SQL = `
		CREATE UNIQUE INDEX IF NOT EXISTS "idx_RegisteredUsers_email" ON "RegisteredUsers"("email");
	`
	q.Exec()
	if !q.Ok() {
		return q.LastError()
	}
	if vars.FeatureUserProfiles {
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
		if !q.Ok() {
			return q.LastError()
		}
	}
}
