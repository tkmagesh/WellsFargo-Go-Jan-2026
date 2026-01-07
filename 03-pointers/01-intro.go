package main

import "fmt"

func main() {
	var no int
	no = 100

	// value -> address
	var noPtr *int
	noPtr = &no // referencing
	fmt.Println("noPtr :", noPtr)

	// address -> value
	x := *noPtr // de-referencing
	fmt.Println(x)

	fmt.Println(no == *(&no))
}
