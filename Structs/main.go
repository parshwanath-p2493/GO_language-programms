// Golang program to show how to
// declare and define the struct

package main

import (
	"fmt"
	"os"
)

func save(v string) {
	data := []byte(v)
	error := os.WriteFile("Person Details", data, 0644)
	if error != nil {
		panic(error)
	}
	fmt.Println("File is Saved.")
}
func main() {
	a := AddNewDetails()
	//var result []string
	result := getAddress(a.name, a.street, a.city, a.state, a.Pincode)
	v := fmt.Sprint(result)
	fmt.Println(v)
	save(v)
}
