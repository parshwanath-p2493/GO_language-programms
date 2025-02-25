package main

import "fmt"

func updatename(X string) {
	X = "John"
}
func updatename2(X string) string {
	X = "John"
	return X
}

func updatename3(X map[int]string) {
	X[1234567890] = "JACK"
}
func main() {
	var t string = "Parshwanath"
	updatename(t)
	fmt.Println(t)
	a := updatename2(t)
	fmt.Println(a)

	var phonenumbers = map[int]string{ //map[KEY_TYPE]VALUE_TYPE
		9338527410: "Parshwanath",
		7418529630: "John",
		8529637410: "Alex",
	}
	updatename3(phonenumbers)
	fmt.Println(phonenumbers)
	for b, c := range phonenumbers {
		fmt.Println(b, "-->", c)
	}

}
