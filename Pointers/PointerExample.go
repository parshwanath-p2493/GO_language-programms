// Golang program to demonstrate the declaration
// and initialization of pointers
package main

import "fmt"

// taking a function with integer
// type pointer as an parameter
func ptf(a *int) {

	// dereferencing
	*a = 748
}
func main() {

	// taking a normal variable
	var x int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &x
	var s *int
	// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
	fmt.Println("Value stored in variable x(*p) = ", *p)
	fmt.Println("Value of variable p = ", &p)
	fmt.Println("s = ", s)

	ptf(p)
	fmt.Printf("The value of x after function call is: %d\n", x)
	ptf(&x)
	fmt.Printf("The value of x after function call is: %d\n", x)

	// taking a variable
	// of integer type
	var V int = 100

	// taking a pointer
	// of integer type
	var pt1 *int = &V

	// taking pointer to
	// pointer to pt1
	// storing the address
	// of pt1 into pt2
	var pt2 **int = &pt1

	fmt.Println("The Value of Variable V is = ", V)
	fmt.Println("Address of variable V is = ", &V)

	fmt.Println("The Value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)

	fmt.Println("The value of pt2 is = ", pt2)

	// Dereferencing the
	// pointer to pointer
	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)

	// double pointer will give the value of variable V
	fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)
}
