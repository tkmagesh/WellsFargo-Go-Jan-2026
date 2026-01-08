/*
V1.0
Add Result : 300
Subtract Result : -100

V2.0
Operation Started...
Add Result : 300
Operation Completed!
Operation Started...
Subtract Result : -100
Operation Completed!
*/
package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	/*
		logAdd := getLogOperation(add)
		logSubtract := getLogOperation(subtract)

		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	add := getLogOperation(add)
	subtract := getLogOperation(subtract)

	add(100, 200)
	subtract(100, 200)
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// V2.0
/*
func logAdd(x, y int) {
	fmt.Println("Operation started...")
	add(x, y)
	fmt.Println("Operation completed!")
}

func logSubtract(x, y int) {
	fmt.Println("Operation started...")
	subtract(x, y)
	fmt.Println("Operation completed!")
}
*/

/*
func logOperation(oper func(int, int), x, y int) {
	fmt.Println("Operation started...")
	oper(x, y)
	fmt.Println("Operation completed!")
}
*/

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Operation started...")
		oper(x, y)
		fmt.Println("Operation completed!")
	}
}
