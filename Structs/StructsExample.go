package main

import (
	"bufio"
	"fmt"
	"os"
)

type Address struct {
	name    string
	street  string
	city    string
	state   string
	Pincode string
}

func getAddress(n, street, c, state string, p string) []string {
	return []string{n, street, c, state, p}
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

	return Address{
		name:    name,
		street:  street,
		city:    city,
		state:   state,
		Pincode: pincode,
	}
}

// func add newDetails()Address{
// 	reader:=bufio.NewReader(os.Stdin)
// }

// Format Printing
func format([]string)
