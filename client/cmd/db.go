package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func mustCreateDB(dbFilePath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		panic(err)
	}
	return db
}
