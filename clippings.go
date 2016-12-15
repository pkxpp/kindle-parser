package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFile(filename string) (*HighlightStorage, *BookStorage, error) {

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	bs := NewBookStorage(GetDb(DbVendor, DbConnString))
	hs := NewHighlightStorage()

	si := 1

	highlight := Highlight{}
	var isErrorOnCurrentIteration bool

	for scanner.Scan() {

		currentString := scanner.Text()

		if len(currentString) > 3 && currentString[0:3] == "===" {
			si = 1
			if !highlight.IsZero() {
				fmt.Println("Adding highlight: ", &highlight, highlight)
				hs.Add(highlight)
				highlight = Highlight{}
				fmt.Println("Highlight after zeeroing: ", &highlight, highlight)
			} // TODO: else log error if highlight is incomplete
			continue
		}

		if isErrorOnCurrentIteration {
			continue
		}

		switch si {
		case 1:
			book, e := CreateBook(currentString) // TODO: this is ugly and probably stupid
			if e != nil {
				log.Printf("Couldn't create a book from string '%s'", currentString)
				isErrorOnCurrentIteration = true
				continue
			}

			book, e = bs.AddIfMissing(book)
			if e != nil {
				log.Printf("Got error while eddit book: %s", e.Error())
			}
			highlight.BookId = book.Id // TODO: return only Id from BookStorage
		case 2:
			highlight.Page, highlight.Location, highlight.Time, _ = parseMetaString(currentString)
		case 4:
			highlight.SetText(currentString)
		}

		si++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &hs, &bs, nil

}
