package main


import (
	"fmt"
//	"regexp"
	"strings"
)

type Book struct {
	Author string
	Title string
}


type BookStorage struct {
	storage map[string][]*Book // should I use list of pointers or pointer on list of books or pointer on list of pointers
}

func NewBookStorage() BookStorage {
	return BookStorage{make(map[string][]*Book)}
}

func (bs *BookStorage) Add(bp *Book) error {
	bs.storage[bp.Author] = append(bs.storage[bp.Author], bp)
	return nil
}

func (bs *BookStorage) AddIfMissing(bp *Book) error {
	if books, ok := bs.storage[bp.Author]; ok {
		for i := range books {
			if bp == books[i] || *bp == *books[i] {
				*bp = *books[i]
				return nil
			}
		}
	}
	return bs.Add(bp)
}

func (bs *BookStorage) Contains(bp *Book) bool {
	if books, ok := bs.storage[bp.Author]; ok {
		for i := range books {
			if bp == books[i] || *bp == *books[i] {
				return true
			}
		}
	}
	return false
}

func (bs *BookStorage) Remove(bp *Book) error {
	if books, ok := bs.storage[bp.Author]; ok {


		for i, book := range books {
			if bp == book || *bp == *book {
				bs.storage[bp.Author] = append(bs.storage[bp.Author][:i], bs.storage[bp.Author][i+1:]...) // TODO: it might panic
			}
		}
	}
	return fmt.Errorf("Storage doesn't contain such book: %v ", bp)
}

func (bs *BookStorage) Len() int {
	res := 0
	for _, s := range bs.storage {
		res += len(s)
	}
	return res
}


func CreateBook(s string) (Book, error) {
	if oi, ci := strings.LastIndex(s, "("), strings.LastIndex(s, ")"); oi != -1 && oi < ci {
		return Book{s[oi+1: ci], strings.Trim(s[:oi], " ")}, nil
	}
	return Book{}, fmt.Errorf("Unknown string format: %s", s)	
}
