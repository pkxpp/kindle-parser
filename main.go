package main

import (
	"fmt"
	"os"
)

func main() {

	hs, bs, e := ParseFile(os.Args[1])

	check(e)

	fmt.Println(hs.Len(), bs.Len())

	highlights := hs.storage

	for i := range highlights {
		fmt.Println(&hs.storage[i], hs.storage[i])
	}

	SaveToDb(hs)

}
