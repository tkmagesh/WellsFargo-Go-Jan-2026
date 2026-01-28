package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateData()
	for range 5 {
		fmt.Println("data :", <-ch)
	}
	fmt.Println("Done")
}

func generateData() <-chan int {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			time.Sleep(2 * time.Second)
			ch <- (i + 1) * 10
		}

	}()
	return ch
}
