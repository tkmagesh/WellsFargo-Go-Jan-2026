package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide be zero error")

func main() {
	// divisor := 7
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("Unknown Error :", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
