package main

import (
	"fmt"
	"time"
)

type Operation func(int, int)
type OperationWrapper func(oper Operation) Operation

func main() {

	add := wrapOperation(add, profileWrapper, logWrapper)
	subtract := wrapOperation(subtract, profileWrapper, logWrapper)

	add(100, 200)
	subtract(100, 200)

}

func wrapOperation(oper Operation, wrapers ...OperationWrapper) Operation {
	for i := len(wrapers) - 1; i >= 0; i-- {
		wrapper := wrapers[i]
		oper = wrapper(oper)
	}
	return oper
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

func logWrapper(oper Operation) Operation {
	return func(x, y int) {
		fmt.Println("Operation started...")
		oper(x, y)
		fmt.Println("Operation completed!")
	}
}

func profileWrapper(oper Operation) Operation {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Printf("Operation execution time : %v\n", elapsed)
	}

}
