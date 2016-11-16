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

	bs := NewBookStorage()
	hs := NewHighlightStorage()

	si := 1

	highlight := Highlight{}	
	
	for scanner.Scan() {

		var bp *Book
		currentString := scanner.Text()

		if len(currentString) > 3 && currentString[0:3] == "===" {
			si = 1
			if ! highlight.IsZero() {
				fmt.Println("Adding highlight: ", highlight)
				hs.Add(&highlight)
				highlight = Highlight{}
			} // TODO: else log error if highlight is incomplete
			continue
		}

		if si == 1 {
			book, e := CreateBook(currentString) // TODO: this is ugly and probably stupid
			bp = &book
			check(e)
			e = bs.AddIfMissing(bp)
			check(e)
			highlight.Book = bp
		} else if si == 2 {
			highlight.Page, highlight.Location, highlight.Time, _ = parseMetaString(currentString)
		} else if si == 4 {
			highlight.Text = currentString
		}
		
		si++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &hs, &bs, nil

}
