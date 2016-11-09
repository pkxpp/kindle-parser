package main


import (
	"fmt"
	"strings"
	"time"
)


//- Your Highlight on Page 5 | Location 64-64 | Added on Tuesday, November 01, 2011, 08:49 AM
// - Your Highlight Location 4407-4408 | Added on Wednesday, November 02, 2011, 12:42 AM


func parseMetaString(s string) (page int, location string, t time.Time, e error) {

	temp := strings.Split(s, "|")

	if len(temp) == 2 {

		location, le := parseLocation(temp[0])
		if le != nil {
			return 0, "", time.Now(), fmt.Errorf("Got error while parsing location in meta string '%s': %v", s, le)
		}
		
		t, te := parseTimeString(temp[1]) 
		if te != nil {
			return 0, "", time.Now(), fmt.Errorf("Got error while parsing time in meta string '%s': %v", s, te)
		}

		return 0, location, *t, nil
	}

	return 0, "wtf", time.Now(), nil
}

func parseLocation(s string) (string, error) {
	return "64-64", nil
}

