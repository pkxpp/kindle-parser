package main

import (
//	"bytes"
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

func ParseFile(filename string) {

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

//	res := make([]string, 1000)
//	var currentRecord bytes.Buffer

	scanner := bufio.NewScanner(file)

	bs := NewBookStorage()
	hs := NewHighlightStorage()

	si := 1

	highlight := Highlight{}	
	
	for scanner.Scan() {

		var bp *Book
		currentString := scanner.Text()

		fmt.Println("\nPrinting string:\n", currentString, "\n")
		
		if len(currentString) > 3 && currentString[0:3] == "===" {
			si = 1
			if ! highlight.IsZero() {
				hs.Add(&highlight)
			}
			continue
		}

		if si == 1 {
			book, e := CreateBook(currentString)
			check(e)
			e = bs.AddIfMissing(&book)
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

}
