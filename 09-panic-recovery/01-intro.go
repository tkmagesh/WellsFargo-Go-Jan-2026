package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("error divide by 0")

func main() {
	defer func() {
		err := recover()
		if errors.Is(err.(error), ErrDivideByZero) {
			fmt.Println("[main - deferred] exiting due to divide by 0 error")
			return
		}
		if err != nil {
			fmt.Println("[main - deferred] exiting due to unknown error :", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	// divisor := 7
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("[main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

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
