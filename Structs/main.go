// Golang program to show how to
// declare and define the struct

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := AddNewDetails()
	//var result []string
	result := getAddress(a.name, a.street, a.city, a.state, a.Pincode)
	fmt.Println(result)
}
func AddNewDetails() Address {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter Street: ")
	street, _ := reader.ReadString('\n')

	fmt.Print("Enter City: ")
	city, _ := reader.ReadString('\n')

	fmt.Print("Enter State: ")
	state, _ := reader.ReadString('\n')

	fmt.Print("Enter Pincode: ")
	pincode, _ := reader.ReadString('\n')

	// Return the struct with trimmed input values
	return Address{
		name:    name,
		street:  street,
		city:    city,
		state:   state,
		Pincode: pincode,
	}

}
