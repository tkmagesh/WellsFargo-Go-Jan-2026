package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
SafeCounter - Custome type that encapsulates the state (count) and "thread safe" behavior (increment)
*/

var count int64

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
	fmt.Println("count :", count)
}

func increment() {
	atomic.AddInt64(&count, 1)
}
