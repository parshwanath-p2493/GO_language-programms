package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hi everyone I'm Parshwanath"
	b := strings.Split(a, " ")
	fmt.Printf("%v \n", b)
	fmt.Println(b[1])
	c := b[1]
	fmt.Println(strings.Index(c, "o"))
	//fmt.Println(a)
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	for index, value := range b {
		fmt.Printf("The index is :%v \t The value is %v\n", index, value)
	}
	fmt.Print("\n")
	for _, value := range b {
		fmt.Printf("The value is %v\n", value)
	}
	fmt.Print("\n")
	for index, _ := range b {
		fmt.Printf("The index is %v\n", index)
	}
}
