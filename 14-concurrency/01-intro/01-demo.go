package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule the execution of f1() through the scheduler to be executed in future
	f2()

	// block the execution of main() so that scheduler looks for other goroutines scheduled and execute them
	time.Sleep(2 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
