package main

import "fmt"

var phonenumbers = map[int]string{ //map[KEY_TYPE]VALUE_TYPE
	9338527410: "Parshwanath",
	7418529630: "John",
	8529637410: "Alex",
}

func main() {
	students := map[string]int{ //map[KEY_TYPE]VALUE_TYPE
		"Parshwanath": 1,
		//	"Parshwanath":1,	--->>Key value must be UNIQUE
		"John":   1,
		"Alex":   2,
		"Denial": 3,
	}
	students["John"] = 25
	fmt.Println(students)
	fmt.Println(students["Alex"])
	for key, value := range students {
		fmt.Println(key, "--->", value)
	}
	fmt.Println(phonenumbers)
	fmt.Println(phonenumbers[8529637410])
	for key, value := range phonenumbers {
		fmt.Printf("%v--->%v\n", key, value)
	}

}
