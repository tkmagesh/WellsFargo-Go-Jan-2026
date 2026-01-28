package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sf atomic.Int64

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		// increment()
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("count :", sf.Load())
}

func increment() {
	sf.Add(1)
}
