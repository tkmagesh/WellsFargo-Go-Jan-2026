package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		add := getLogOperation(add)
		subtract := getLogOperation(subtract)
	*/

	add := getProfiledOperation(getLogOperation(add))
	subtract := getProfiledOperation(getLogOperation(subtract))

	add(100, 200)
	subtract(100, 200)

}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Operation started...")
		oper(x, y)
		fmt.Println("Operation completed!")
	}
}

func getProfiledOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Printf("Operation execution time : %v\n", elapsed)
	}

}
