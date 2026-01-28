package main

import (
	"fmt"
	"time"
)

// variable used for signalinng (but follows communicate by sharing memory)
var shouldStop bool = false

func main() {
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		shouldStop = true
	}()
	doneCh := doSomething()
	<-doneCh
}

func doSomething() <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		for i := 1; ; i++ {
			if !shouldStop {
				time.Sleep(200 * time.Millisecond)
				fmt.Println(i)
			} else {
				break
			}
		}
		close(doneCh)
	}()
	return doneCh
}
