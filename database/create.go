package database

import (
	"database/sql"

	"github.com/suite911/query911/query"
)

const defaultNow = `DEFAULT(CAST(strftime('%s', 'now') AS INTEGER))`

func Create(db *sql.DB) {
	q := query.Query{DB: db}
	q.SQL = `
		CREATE TABLE IF NOT EXISTS "RegisteredUsers" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"email" TEXT NOT NULL UNIQUE,
			"pw" BLOB, -- specifically can be NULL
			"credits" INTEGER NOT NULL DEFAULT(0),
			"ts" INTEGER NOT NULL ` + defaultNow + `
		);
		CREATE UNIQUE INDEX IF NOT EXISTS "idx_RegisteredUsers_email" ON "RegisteredUsers"("email");
	`
}
