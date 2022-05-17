package main

import "fmt"

//here We have Defined Two Ints
//indicating that There are multiple Return Statements
func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
