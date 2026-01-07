package main

import "fmt"

func main() {

	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
		fmt.Println("no = ", no)
	*/

	if no := 21; no%2 == 0 { // "no" is scoped within the if-else block (not accessible outside the if-else block)
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}
	fmt.Println("no = ", no)
}
