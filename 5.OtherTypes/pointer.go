package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

//ZeroPtr will take an int Pointer
func zeroptr(iptr *int) {
	*iptr = 0

}
func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroVale:", i)
	//printing value of i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	//printing Value of Pointer
	fmt.Println("pointer:", &i)

}
