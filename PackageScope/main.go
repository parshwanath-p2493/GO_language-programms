package main

/** -----PackageScope(Folder)
|--main.go
|--branch.go
*/

import (
	"fmt"
	"sort"

	"rsc.io/quote"
)

var score int = 50

func main() {
	//var score int = 50--------- Its invalid , we need to define this at root level

	fmt.Println(quote.Glass())
	HelloWorld("Parshwanth")
	sort.Ints(a)
	for _, v := range a {
		fmt.Println(v)
	}
	Score(score)
}
