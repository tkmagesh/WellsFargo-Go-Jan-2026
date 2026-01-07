package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hi!")
	}()

	func(userName string) {
		fmt.Printf("Hello, %s!\n", userName)
	}("Suresh")

	q, r := func(multiplier, divisor int) (quotient, remainder int) {
		quotient = multiplier / divisor
		remainder = multiplier % divisor
		return
	}(100, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", q, r)
}
