package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello, %s!\n", userName)
	}
	sayHello("Magesh")

	var divide func(int, int) (int, int)
	divide = func(multiplier, divisor int) (quotient, remainder int) {
		quotient = multiplier / divisor
		remainder = multiplier % divisor
		return
	}
	quotient, remainder := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remaider = %d\n", quotient, remainder)

}
