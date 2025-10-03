// Function returning two values----
package main

import (
	"fmt"
	"strings"
)

func main() {
	value1, value2 := getTwoValues("Parshwanath mahaveer paramagond")
	fmt.Println(value1, value2)
	value3, value4 := getTwoValues("u cant see me")
	fmt.Println(value3, value4)
}

func getTwoValues(n string) (string, string) {
	a := strings.ToUpper(n)
	b := strings.Split(a, " ")
	var initials []string
	for _, value := range b {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}
