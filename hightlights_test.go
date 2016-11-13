package main

import (
	"fmt"
	"time"
	"testing"
)

var books = []Book {
	{"Iain Banks", "Excession"},
	{"Iain Banks", "State of the art"},
	{"Bertrand Russell", "A History of Western Philosophy"},
	{"Robert Martin", "Clean Code"},			
}

var highlights = []Highlight {
	{&books[0], "some word", "30-31", time.Date(2016, 4, 19, 6, 36, 12, 0, time.UTC)},
	{&books[0], "exception", "340-340", time.Date(2016, 5, 20, 16, 20, 12, 0, time.UTC)},
	{&books[0], " Entertain", "30-31", time.Date(2016, 4, 19, 6, 33, 2, 0, time.UTC)},
	{&books[1], "Some very long string", "30-31", time.Date(2015, 9, 19, 13, 36, 17, 0, time.UTC)},
	{&books[1], "exception", "1-3", time.Date(2016, 5, 20, 16, 20, 12, 0, time.UTC)},
	{&books[3], "WTF???????????????", "30-31", time.Date(2015, 10, 19, 7, 36, 23, 0, time.UTC)},
}


func TestHighlightStorage (t *testing.T) {

	hs := NewHighlightStorage()

	for _, h := range highlights {
		hs.Add(&h)
	}

//	hs.Add(&highlights[0])

	fmt.Println(hs)
	fmt.Println( highlights)

	checkLen(hs.Len(), len(highlights), "HighlightStorage", t)
}


func checkLen(actual, expected int, name string, t *testing.T) {
	if actual != expected {
		t.Errorf("Wrong Len of %s. Actual: %d. Expected: %d", name, actual, expected)
	}
}




