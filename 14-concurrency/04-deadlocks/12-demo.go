package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for range 5 {
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
	fmt.Println("Done!")
}

func fn() {
	fmt.Println("fn invoked")
}
