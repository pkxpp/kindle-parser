package main

import (
	"fmt"
	"os"
)


func main() {

	hs, bs, e := ParseFile(os.Args[1])

	check(e)
	
	fmt.Println(hs.Len(), bs.Len())

	highlights := hs.hs

	for i := range highlights {
		fmt.Println(&hs.hs[i], hs.hs[i])
	}


	books:= bs.storage

	for i := range books {
		fmt.Println(&books[i], books[i])
	}
	


	saveBooks(bs)

//	SaveToDb(hs)

}
