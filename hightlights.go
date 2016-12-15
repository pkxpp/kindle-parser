package main

import (
	//	"bytes"
	//	"bufio"
	"fmt"
	"log"
	//	"os"
	"strings"
	"time"
	//	"github.com/davecgh/go-spew/spew"
)

type Highlight struct {
	BookId   uint
	Text     string
	Page     int
	Location string
	Time     time.Time
}

func (h *Highlight) IsZero() bool {
	return h.BookId == 0 && h.Text == "" && h.Page == 0 && h.Location == "" && h.Time.IsZero()
}

func (h *Highlight) SetText(text string) {
	h.Text = strings.Trim(text, " ,.'\"“’ ”")
}

type HighlightStorage struct {
	storage []Highlight
	byText  map[string][]uint
}

func NewHighlightStorage() HighlightStorage { // TODO: default argument
	return HighlightStorage{
		make([]Highlight, 0, 20),
		make(map[string][]uint),
	}
}

func (hs *HighlightStorage) Add(h Highlight) error {
	if hs.Contains(h) {
		return fmt.Errorf("Highlight already exists: ", h)
	}

	//	fmt.Println("From hs.Add: ", &h, h)
	hs.storage = append(hs.storage, h)
	//	fmt.Println("Hs after appending: ", hs.storage)

	index := len(hs.storage) - 1
	if index < 0 {
		log.Fatalf("Index less than zero and it can lead to strage shit. Index: %d", index)
	}

	hs.byText[h.Text] = append(hs.byText[h.Text], uint(index))

	//	fmt.Println("hs.ByText, hs.byBook: ", hs.byText, hs.byBook)
	return nil
}

func (hs *HighlightStorage) Contains(h Highlight) bool {
	if _, ok := hs.byText[h.Text]; !ok {
		return false
	}
	for _, index := range hs.byText[h.Text] {
		if hs.storage[index] == h { // TODO: check that Book field equeal!
			// TODO: think about Id!!
			return true
		}
	}
	return false
}

func (hs *HighlightStorage) Len() int {
	return len(hs.storage)
}

func (hs *HighlightStorage) GetByText(t string) ([]Highlight, error) {
	if indexes, ok := hs.byText[t]; ok {
		res := make([]Highlight, 0, len(indexes))
		for _, index := range indexes {
			res = append(res, hs.storage[index])
		}
		return res, nil
	}
	return nil, fmt.Errorf("Highlight with such text doesn't exist (%s)", t)
}

func ParseHighlight(s string) {

}

func ParseClippingsFile(s string) HighlightStorage {
	return HighlightStorage{}
}
