package main

import (
	"fmt"
	"strings"
	"regexp"
	"time"
)

type timeString struct  {
	Regexp string
	String string
}

var knownTimeStrings = []timeString {
	{`[\d]{1,2}[\s][\w]+[\s][\d]{4}`, `2 January 2006 15:04:05`},
	{`[\d]{1,2}[\s][\w]+[\s][\d]{2}`, `2 January 06 15:04:05`},
	{`[\w]+[\s][\d]{1,2},[\s][\d]{4}`, `January 2, 2006, 03:04 PM`},
}

func parseTimeString(s string) time.Time {

	if len(s) < 10 {
		panic(fmt.Sprintf("Passed string is too short: %s", s))
	}

	s = s[strings.Index(s, ",") + 2:]

	fmt.Println(s)

	for _, ts := range knownTimeStrings {

		if m, e := regexp.MatchString(ts.Regexp, s); m {
			
			if res, err := time.Parse(ts.String, s); err != nil {
				panic(err)
			} else {
				return res
			}
		} else if e != nil{
			panic(e)
		}
		
	}

	panic(fmt.Sprintf("Unknown string given. %s", s))
}
