package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbVendor, DbConnString = "sqlite3", "db.sqlite"

var db sql.DB

func GetDb(DbVendor, DbConnString string) *sql.DB {
	db, e := sql.Open(DbVendor, DbConnString)
	db.Exec("PRAGMA foreign_keys = ON;")
	if e != nil {
		log.Fatal(e)
	}
	return db
	
}

func SaveToDb(hs *HighlightStorage) {

	db, e := sql.Open(DbVendor, DbConnString)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close() // TODO: defer in init()

	highlights := hs.storage

	for i, highlight := range highlights {

		stmt, err := db.Prepare(`
INSERT INTO highlights(text, page, location, time, book_id) VALUES(?, ?, ?, ?, ?)
`)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(highlight.Text, highlight.Page, highlight.Location, highlight.Time, highlight.BookId)
		if err != nil {
			log.Fatal(err)
		}

		lastId, err := res.LastInsertId()

		fmt.Println("Last id", lastId, i)
		if err != nil {
			log.Fatal(err)
		}

		//		highlights[i].Id = uint(lastId)

	}

}
