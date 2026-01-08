package main

import (
	"errors"
	"fmt"
)

func main() {
	divisor := 7
	// divisor := 0
	q, r, err := divide(100, divisor)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

/*
func divide(multiplier, divisor int) (int, int, error) {
	var quotient, remainder int
	if divisor == 0 {
		err := errors.New("do not attempt to divide by zero")
		return 0, 0, err
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return quotient, remainder, nil
}
*/

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		err = errors.New("do not attempt to divide by zero")
		return
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
