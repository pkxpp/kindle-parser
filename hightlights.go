package main

import (
//	"bytes"
//	"bufio"
	"fmt"
//	"log"
//	"os"
	"time"
//	"github.com/davecgh/go-spew/spew"
)


type Highlight struct {
	Book *Book
	Text string
	Page int
	Location string
	Time time.Time
}

func (h *Highlight) IsZero() bool {
	return h.Book == nil && h.Text == "" && h.Page == 0 && h.Location == "" && h.Time.IsZero()
}

type HighlightStorage struct {
	hs []*Highlight
	byText map[string][]*Highlight
	byBook map[Book][]*Highlight
}

func NewHighlightStorage() HighlightStorage {
	return HighlightStorage{
		make([]*Highlight, 0, 20),
		make(map[string][]*Highlight),
		make(map[Book][]*Highlight),
	}
}

func (hs *HighlightStorage) Add(h *Highlight) error {
	if hs.Contains(h) {
		return fmt.Errorf("Highlight already exists: ", h)
	}
	hs.hs = append(hs.hs, h) 
	hs.byText[h.Text] = append(hs.byText[h.Text], h)
	hs.byBook[*h.Book] = append(hs.byBook[*h.Book], h)
	return nil
}

func (hs *HighlightStorage) Contains(h *Highlight) bool {
	if _, ok := hs.byText[h.Text]; !ok {
		return false
	}
	for _, hp := range hs.byText[h.Text] {
		if hp == h || *hp == *h { // TODO: check that Book field equeal! 
			return true 
		}
	}
	return false
}

func (hs *HighlightStorage) Len() int {
	return len(hs.hs)
}

func (hs *HighlightStorage) GetByText(t string) ([]*Highlight, error) {
	if res, ok := hs.byText[t]; ok {
		return res, nil
	}
	return []*Highlight{}, fmt.Errorf("Highlight with such text doesn't exist (%s)", t)	
}


func ParseHighlight(s string, ) {


}


func ParseClippingsFile(s string) HighlightStorage {
	return HighlightStorage{}
}
