package main

import (
	"fmt"
	"time"
)

// signaling using "share memory by communicating"

func main() {
	stopCh := time.After(5 * time.Second)
	doneCh := doSomething(stopCh)
	<-doneCh
}

func doSomething(stopCh <-chan time.Time) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-stopCh:
				fmt.Println("Stop signal received")
				break LOOP
			default:
				time.Sleep(200 * time.Millisecond)
				fmt.Println(i)
			}
		}
		close(doneCh)
	}()
	return doneCh
}
