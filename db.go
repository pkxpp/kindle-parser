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


	stmt, err := db.Prepare("INSERT INTO highlights(text, ) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("Dolly")
	if err != nil {
		log.Fatal(err)
	}
	


	var t string
	
	e := db.QueryRow("select t from test;").Scan(&t)

	fmt.Println(t, e)

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
