package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	fmt.Print(getGreetMsg("Magesh", "Kuppan"))

	fmt.Println(divide(100, 7))

	// using the results of the divide function
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remaider = %d\n", quotient, remainder)
	*/

	// using ONLY the quotient
	/*
		quotient, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
	*/

}

func sayHi() {
	fmt.Println("Hi!")
}

func sayHello(userName string) {
	fmt.Printf("Hello, %s!\n", userName)
}

/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

func getGreetMsg(firstName, lastName string) string {
	s := fmt.Sprintf("Hi %s %s, Have a good day!\n", firstName, lastName)
	return s
}

/*
func divide(multiplier, divisor int) (int, int) {
	var quotient, remainder int
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return quotient, remainder
}
*/

// named results
func divide(multiplier, divisor int) (quotient, remainder int) {
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
