package main


import (
	"fmt"
	//	"regexp"
	"strings"
)

type Book struct {
	Id int	
	Author string
	Title string
}

func NewBook(author, title string) (Book) {
	return Book{0, author, title}
}

type BookStorage struct {
	storage []Book
	byAuthor map[string][]uint // should I use list of pointers or pointer on list of books or pointer on list of pointers
}

func NewBookStorage() BookStorage {
	return BookStorage{make([]Book, 0), make(map[string][]uint)}
}

func (bs *BookStorage) Add(bp *Book) error {
	bs.storage = append(bs.storage, *bp)
	bs.byAuthor[bp.Author] = append(bs.byAuthor[bp.Author], uint(len(bs.storage)-1))
	return nil
}

func (bs *BookStorage) AddIfMissing(bp *Book) error { // TODO: now it doesn't change pointer like before | It's too complicated and error-prone
	if indexes, ok := bs.byAuthor[bp.Author]; ok {
		for _, i := range indexes {
			if *bp == bs.storage[i] {
				bp = &bs.storage[i] // TODO: this is pointles
				return nil
			}
		}
	}
	return bs.Add(bp) 
}

func (bs *BookStorage) Contains(bp *Book) bool {
	if indexes, ok := bs.byAuthor[bp.Author]; ok {
		for _, i := range indexes {
			if *bp == bs.storage[i] {
				return true
			}
		}
	}
	return false
}

func (bs *BookStorage) Remove(bp *Book) error {
	if indexes, ok := bs.byAuthor[bp.Author]; ok {
		for i := range indexes {
			if *bp == bs.storage[indexes[i]] {
				bs.storage = append(bs.storage[:i], bs.storage[i+1:]...) // TODO: it might be very costly, consider list usage
				bs.byAuthor[bp.Author] = append(bs.byAuthor[bp.Author][:i], bs.byAuthor[bp.Author][i+1:]...) // TODO: it might panic
			}
		}
	}
	return fmt.Errorf("Storage doesn't contain such book: %v ", bp)
}

func (bs *BookStorage) Len() int {
	res := 0
	for _, s := range bs.byAuthor {
		res += len(s)
	}
	return res
}


func CreateBook(s string) (Book, error) {
	if oi, ci := strings.LastIndex(s, "("), strings.LastIndex(s, ")"); oi != -1 && oi < ci {
		return Book{0, s[oi+1: ci], strings.Trim(s[:oi], " ")}, nil
	}
	return Book{0, "", strings.Trim(s, " ")}, nil
}
