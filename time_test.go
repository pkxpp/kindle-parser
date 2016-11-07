package main

import (
	"fmt"
	"testing"
	"time"
)

var parseTimeStringTests = []struct {
	s string
	expected time.Time
} {
	{
		"Added on Monday, 31 October 2016 07:27:46",
		time.Date(2016, 10, 31, 7, 27, 46, 0, time.UTC),
	},
	{
		"Added on Saturday, 11 July 15 02:25:50",
		time.Date(2015, 07, 11, 2, 25, 50, 0, time.UTC),		
	},
	{
		"Added on Monday, 30 October 2016 07:27:46",
		time.Date(2016, 10, 30, 7, 27, 46, 0, time.UTC),
	},
	{
		"Added on Friday, March 28, 2014, 10:13 PM",
		time.Date(2014, 3, 28, 22, 13, 0, 0, time.UTC),		
	},
}

func TestParseTimeString(t *testing.T) {

	fmt.Println("running TestParseTimeString")

	for _, pair := range parseTimeStringTests {
		
		res := parseTimeString(pair.s)

		if !res.Equal(pair.expected) {
			t.Error("Date not equal! string: '", pair.s, "'| res: ", res.Local(), "| expected: ", pair.expected.Local())
		}
	}

}
