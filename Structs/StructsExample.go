package main

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

// func add newDetails()Address{
// 	reader:=bufio.NewReader(os.Stdin)
// }
/*
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r. ReadString('\n')

	return strings. TrimSpace(input), err

	func createBill() bill {
	reader := bufio. NewReader(os.Stdin)

	:= getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
*/
