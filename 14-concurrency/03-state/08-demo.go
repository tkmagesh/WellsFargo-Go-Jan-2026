package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
SafeCounter - Custome type that encapsulates the state (count) and "thread safe" behavior (increment)
*/
type SafeCounter struct {
	sync.Mutex
	count int
}

func (sf *SafeCounter) Add(delta int) {
	sf.Lock()
	{
		sf.Add(1)
	}
	sf.Unlock()
}

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
