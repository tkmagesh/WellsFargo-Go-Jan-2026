package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

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
	mutex.Lock()
	{
		count += 1
	}
	mutex.Unlock()
}
