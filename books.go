package main

import (
	"database/sql"
	"fmt"
	//	"regexp"
	"strings"
)

type Book struct {
	Id     uint
	Author string
	Title  string
}

func NewBook(author, title string) Book {
	return Book{0, author, title}
}

func (b *Book) Equals(other Book) bool {
	if b.Author==other.Author && b.Title==other.Title {
		return true
	}
	return false 
}

type BookStorage struct {
	byId map[uint]Book
	byAuthor map[string][]uint // should I use list of pointers or pointer on list of books or pointer on list of pointers
	db *sql.DB
}

type NoSuchBook struct{
	book Book
}

func (e NoSuchBook) Error() string {
	return fmt.Sprintf("No such book in storage: %s", e.book )
}

func NewBookStorage(db *sql.DB) BookStorage {
	return BookStorage{make(map[uint]Book), make(map[string][]uint), db}
}

func (bs *BookStorage) Books() map[uint]Book {
	return bs.byId
}

func (bs *BookStorage) Add(b Book) (Book, error) {
	b, err := bs.store(b)

	fmt.Printf("From bs.Add 1 | b: %v | err: %v\n", b, err)
	if err != nil {
		return b, err // TODO: maybe better return Book{} ? 
	}
	bs.byId[b.Id] = b
	bs.byAuthor[b.Author] = append(bs.byAuthor[b.Author], b.Id)
	return b, err
}

func (bs *BookStorage) store(b Book) (Book, error) {
	stmt, err := bs.db.Prepare("INSERT INTO books(author, title) VALUES(?, ?)")
	if err != nil {
		return b, err
	}
	res, err := stmt.Exec(b.Author, b.Title)
	if err != nil {
		return b, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return b, err
	}
	b.Id = uint(lastId)
	return b, nil
}

func (bs *BookStorage) AddIfMissing(b Book) (Book, error) { // TODO: now it doesn't change pointer like before | It's too complicated and error-prone
	id, err := bs.GetIdByBook(b)
	_, okError := err.(NoSuchBook)
	if err == nil && id != 0 {
		return bs.byId[id], nil
	} else if !okError {
		return b, err
	}

	return bs.Add(b)
}

func (bs *BookStorage) GetIdByBook(b Book) (id uint, err error) {
	if indexes, ok := bs.byAuthor[b.Author]; ok {
		for _, i := range indexes {
			if book := bs.byId[i]; book.Equals(b) {
				if book.Id == 0 {
					// TODO: it must never happen, I probably shouldn't care
					return 0, fmt.Errorf("Book (%s) in storage but doesn't have valid Id. bs.storage id: %d", bs.byId[i], i)
				}
				return book.Id, nil
			}
		}
	}
	return 0, NoSuchBook{b}
}


func (bs *BookStorage) Contains(b Book) bool {
	return false
}

func (bs *BookStorage) Remove(b Book) error {
	return nil
}

func (bs *BookStorage) Len() int {
	return len(bs.byId)
}

func CreateBook(s string) (Book, error) {
	if oi, ci := strings.LastIndex(s, "("), strings.LastIndex(s, ")"); oi != -1 && oi < ci {
		return Book{0, s[oi+1 : ci], strings.Trim(s[:oi], " ")}, nil
	}
	return Book{0, "", strings.Trim(s, " ")}, nil
}
