package main

import (
	"bytes"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)


type Book struct {
	Author string
	Title string
}

type Highlight struct {
	Text string
	Location string
	Time time.Time
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}




func main() {

	file, err := os.Open("examples/ex3.txt")
	check(err)

	defer file.Close()

	res := make([]string, 1000)
	var currentRecord bytes.Buffer

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		
		currentString := scanner.Text()
		
		if len(currentString) > 3 && currentString[0:3] == "===" {
			res = append(res, currentRecord.String())
			currentRecord.Reset()
			continue
		}



		currentRecord.WriteString(currentString)				
			

	}

	fmt.Println("%v", res)	

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
