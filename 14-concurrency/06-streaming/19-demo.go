package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := generateData()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(1 * time.Second)
			fmt.Println("data :", data)
			continue
		}
		break
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
