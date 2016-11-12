package main

import (
	"fmt"
	"time"
	"testing"
)


func TestHightlights (t *testing.T) {

	hc := Highlights{make([]*Highlight, 0, 10), make(map[string][]*Highlight)}

	b := Book{"Iain Banks", "Excession"}
	
	h := Highlight{&b, "garbled", "237-237", time.Date(2016, 04, 19, 6, 36, 12, 0, time.UTC)}

	if e := hc.Add(&h); e != nil {
		t.Errorf("Couldn't Add Hightlight", h)
	}

	if hc.Contains(&h) != true {
		t.Errorf("Contains returns wrong result", h)
	}

	fmt.Println(hc)
}



