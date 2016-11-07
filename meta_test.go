package main

import (
	"fmt"
	"testing"
)


var metaTests = []struct {
	S string
	Page int
	Location string
} {
	{
		`- Your Highlight on Page 5 | Location 64-64 | Added on Tuesday, November 01, 2011, 08:49 AM`,
		5,
		`64-64`,
	},
	{
		`- Your Highlight Location 4407-4408 | Added on Wednesday, November 02, 2011, 12:42 AM`,
		0,
		`4407-4408`,
	},
}


func TestParseMetaString(t *testing.T) {

	for _, data := range metaTests {

		page, location, _ := parseMetaString(data.S)

		if page != data.Page {
			t.Error("Page not equal. Expected: ", data.Page, " | Got: ", page)
		}

		if location != data.Location {
			t.Error("Location not equal. Expected: ", data.Location, " | Got: ", location)
		}
	}

}
