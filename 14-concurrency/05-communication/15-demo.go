package main

import "fmt"

func main() {
	/*
		ch := make(chan int)
		data := <-ch
		ch <- 100
		fmt.Println("data :", data)
	*/

	/*
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println("data :", data)
	*/

	// (1) create a channel
	ch := make(chan int)

	// (2) schedule the function to execute in future (goroutine)
	go func() {
		// (4) send the data through the open door (non-blocking)
		ch <- 100

		// (5) goroutines completes
	}()

	// (3) Open the door (channel) and wait (block) for the data to arrive
	data := <-ch
	// (6) receive the data and unblock the receive operation

	// (7) print the data
	fmt.Println("data :", data)
}
