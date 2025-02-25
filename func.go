package main

import (
	"fmt"
	"math"
)

func studentName(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}
func area(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	// Hello("Parshwanath")
	// goodbye("Parshwanath")
	studentName([]string{"parshwanath", "pbcckk", "xyz", "lmnopqrstuvwxyz"}, Hello)
	studentName([]string{"parshwanath", "pbcckk", "xyz", "lmnopqrstuvwxyz"}, goodbye)

	fmt.Println(area(123))
	fmt.Println(area(1))

}

func Hello(n string) {
	fmt.Printf(" HI Hello!! my name is %v\n", n)
}
func goodbye(n string) {
	fmt.Printf(" This is  goodbye......... from ur %v\n", n)
}

}