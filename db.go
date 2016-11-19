package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


var DbVendor, DbConnString = "sqlite3", "db.sqlite"


var db sql.DB

func init() {
}


func SaveToDb(hs *HighlightStorage) {

	db, e := sql.Open(DbVendor, DbConnString)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close() // TODO: defer in init()
	
	highlights := hs.hs

	for i, highlight := range highlights {

		stmt, err := db.Prepare(`
INSERT INTO highlights(text, page, location, time, book_id) VALUES(?, ?, ?, ?, ?)
`)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(highlight, *highlight)
		fmt.Println(highlights[i], *highlights[i])
		fmt.Println(highlights[i:i+10], highlights[i:i+10])
		
		res, err := stmt.Exec(highlight.Text, highlight.Page, highlight.Location, highlight.Time, 1)
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


func saveBooks(bs *BookStorage) (error) {

	db, e := sql.Open(DbVendor, DbConnString)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close() // TODO: defer in init()
	
	books := *bs.Books()

	for i := range books {
		stmt, err := db.Prepare("INSERT INTO books(author, title) VALUES(?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(books[i].Author, books[i].Title)
		if err != nil {
			log.Fatal(err)
		}

		lastId, err := res.LastInsertId()

		fmt.Println("Last id", lastId)
		if err != nil {
			log.Fatal(err)
		}

		books[i].Id = uint(lastId)
	}
	
	
	return nil
}
