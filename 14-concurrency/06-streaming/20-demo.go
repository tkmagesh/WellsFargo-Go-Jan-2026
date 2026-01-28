package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := generateData()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("data :", data)
	}
	fmt.Println("Done")
}

func generateData() <-chan int {
	ch := make(chan int)
	count := rand.Intn(20)
	fmt.Printf("[generateData] - about to generate %d numbers\n", count)
	go func() {
		for i := range count {
			ch <- (i + 1) * 10
		}
		// signal the end of the data stream
		close(ch)
	}()
	return ch
}
