package main

import (
	"database/sql"
	"io/ioutil"
	_ "github.com/mattn/go-sqlite3"
)


func getTestDB() *sql.DB {

	DbVendor, DbConnString := "sqlite3", ":memory"

	db, e := sql.Open(DbVendor, DbConnString)
	if e != nil {
		panic(e)
	}

	sql, e := ioutil.ReadFile("db.sql")
	if e != nil {
		panic(e)
	}

	_, e = db.Exec(string(sql))
	if e != nil {
		panic(e)
	}

// 	defer db.Close() // TODO: do I need to close it? Where?

	return db
}
