package main

import "fmt"

var nameFive int = 7474185645642586978

// nameSix :=121030 // command-line-arguments
//	.\first.go:6:1: syntax error: non-declaration statement outside function body

func main() {
	var firstName string = "Parshwanath "
	var secondName = "Paramagond"
	var thirdName string
	fmt.Println(firstName, secondName, thirdName)
	//-------------------------------------------------------------------------------------------------
	firstName = "ASDFG"
	thirdName = "PBCCKK"
	fmt.Println(firstName, secondName, thirdName, nameFive)
	nameFour := 123
	nameSix := 121030
	fmt.Print(nameFour, "\n", nameSix) //short variablr decleration
	//we cant let the declared variable unused it must be used to print somewhere
	fmt.Print(nameFour, "\n", nameSix)
	fmt.Printf("\nMy name is %v and my number is %v \n \n", firstName, nameFour)

	//------------------------------------------------------------------------------------------------------
	// %_ are the format specifiers
	// %v for variables %q for the quotes ------  % T is used to print the type of variable-----%0.1f for float type modification
	a := 14
	b := 50.2246
	fmt.Printf("The type of var a is --%T \n", a)
	fmt.Printf("The value of b is -- %0.1f \n ", b)

	//------------------------------------------------------------------------------------------------------------------
	// Sprint is used to save the string/output ...example
	var abc = fmt.Sprintf("My name is %v and my number is %v \n", firstName, nameFour)
	fmt.Println("The saved string is :- ", abc)

	//--------------------------------------------------------------------------------------------------------------
	// ARRAYS & SLICES

	var num [4]int = [4]int{10, 20, 40, 60}
	fmt.Println(num, len(num))
	var num2 = [4]int{10, 40, 60} // OUTPUT --[10 40 60 0] 4
	fmt.Println(num2, len(num2))

	names := [4]string{"plm", "PBCCKK", "PT"}
	fmt.Println(names, len(names)) // OUTPUT-- [plm PBCCKK PT ] 4

	//SLICES
	var nums3 = []int{100, 500, 400, 900}
	nums3 = append(nums3, 741)
	fmt.Println(nums3, len(nums3))

	//RANGES OF SLICES
	range1 := num[1:3]
	range2 := nums3[:]
	range3 := names[1:]
	fmt.Println(range1, "\n", range2, "\n", range3)

}

//https://pkg.go.dev/std
//https://go.dev/ref/spec#Boolean_types
