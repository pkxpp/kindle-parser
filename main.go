package main

import (
	"fmt"
	"os"
)


func main() {

	hs, bs, e := ParseFile(os.Args[1])

	check(e)
	
	fmt.Println(hs.Len(), bs.Len())


	saveBooks(bs)

}
