package main

import (
	"fmt"
)

// share memory by communicating

func main() {
	// var ch chan int = make(chan int)
	// var ch = make(chan int)
	ch := make(chan int)

	go add(100, 200, ch)
	// receive the result from the channel
	result := <-ch

	fmt.Println("Result :", result)
}

func add(x, y int, ch chan int) {
	result := x + y
	// send the result through the channel
	ch <- result

}
