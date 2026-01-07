package main

import "fmt"

func main() {
	/*
		var username string
		fmt.Println("Enter your name :")
		fmt.Scanln(&username)
		fmt.Println("Entered name :", username)
	*/

	var firstName, lastName string
	fmt.Println("Enter your complete name:")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Println(lastName, ", ", firstName)
	fmt.Printf("%d, %s\n", lastName, firstName)
	fmt.

}
