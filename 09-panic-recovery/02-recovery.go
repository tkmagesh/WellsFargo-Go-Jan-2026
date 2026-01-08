package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("error divide by 0")

func main() {
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		if q, r, err := divideAdapter(100, divisor); err == nil {
			fmt.Printf("[main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			break
		} else {
			fmt.Printf("Error : %v, \n Try again!\n", err)
		}
	}
}

// convert the panic into an error
func divideAdapter(multiplier, divisor int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(multiplier, divisor)
	return
}

// 3rd party library
func divide(multiplier, divisor int) (quotient, remainder int) {
	defer func() {
		if err := recover(); err != nil {
			divideErr := fmt.Errorf("divide: multiplier : %d, divisor : %d, Error : %w, source : %w", multiplier, divisor, ErrDivideByZero, err)
			panic(divideErr)
		}
	}()
	fmt.Println("[divide] attempting to find quotient")
	/* if divisor == 0 {
		panic(ErrDivideByZero)
	} */
	quotient = multiplier / divisor
	fmt.Println("[divide] attempting to find remainder")
	remainder = multiplier % divisor
	return
}
