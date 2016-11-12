package main

import (
//	"bytes"
//	"bufio"
	"fmt"
//	"log"
//	"os"
	"time"
	
)


type Highlight struct {
	Book *Book
	Text string
	Location string
	Time time.Time
}

type Highlights struct {
	hc []*Highlight
	hf map[string][]*Highlight
}

func (hc *Highlights) Add(h *Highlight) error {
	if hc.Contains(h) {
		return fmt.Errorf("Highlight already exists: ", h)
	}
	hc.hc = append(hc.hc, h) 
	hc.hf[h.Text] = append(hc.hf[h.Text], h)
	return nil
}

func (hc *Highlights) Contains(h *Highlight) bool {
	if _, ok := hc.hf[h.Text]; !ok {
		return false
	}
	for _, hp := range hc.hf[h.Text] {
		if hp == h || *hp == *h { // TODO: check that Book field equeal! 
			return true 
		}
	}
	return false
}

func (hc *Highlights) GetByText(t string) ([]*Highlight, error) {
	if res, ok := hc.hf[t]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("Highlights with such text doesn't exist (%s)", t)	
}


func ParseHighlight(s string, ) {


}


func ParseClippingsFile(s string) Hightlights {
	return nil
}
