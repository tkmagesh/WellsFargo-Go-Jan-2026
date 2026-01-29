package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// signaling using "context"

func main() {
	rootCtx := context.Background()
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")
	cancelCtx, cancel := context.WithCancel(valCtx)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	doneCh := doSomething(cancelCtx)
	<-doneCh
}

func doSomething(ctx context.Context) <-chan struct{} {

	fmt.Printf("[doSomething] root-key : %v\n", ctx.Value("root-key"))
	// ctx = context.WithValue(ctx, "root-key", "new-root-value")
	doneCh := make(chan struct{})
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		fibCancelCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		go genFib(fibCancelCtx, wg)

		wg.Add(1)
		primeCancelCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		go genPrimes(primeCancelCtx, wg)
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("[doSomething] Stop signal received")
				break LOOP
			default:
				time.Sleep(200 * time.Millisecond)
				fmt.Println("doSomething : ", i)
			}
		}
		wg.Wait()
		close(doneCh)
	}()
	return doneCh
}

func genFib(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[genFib] root-key : %v\n", ctx.Value("root-key"))
LOOP:
	for x, y := 0, 1; ; x, y = y, x+y {
		select {
		case <-ctx.Done():
			fmt.Println("[genFib] cancellation signal received")
			break LOOP
		default:
			fmt.Printf("Fib : %d\n", x)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func genPrimes(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[genPrimes] root-key : %v\n", ctx.Value("root-key"))
LOOP:
	for no := 2; ; no++ {
		select {
		case <-ctx.Done():
			fmt.Println("[genPrimes] cancellation signal received")
			break LOOP
		default:
			if isPrime(no) {
				fmt.Printf("Prime : %d\n", no)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
