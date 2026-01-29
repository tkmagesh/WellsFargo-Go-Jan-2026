package main

import (
	"context"
	"fmt"
	"time"
)

// signaling using "context"

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	fmt.Println("Will stop after 15 secs, but Hit ENTER to stop before 15 secs...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	doneCh := doSomething(cancelCtx)
	<-doneCh
}

func doSomething(ctx context.Context) <-chan struct{} {

	doneCh := make(chan struct{})
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
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
