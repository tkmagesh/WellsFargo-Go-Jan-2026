package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Refactor the below to follow "Share memory by communicating"
*/

func main() {
	var start, end int
	fmt.Println("Enter the start and end")
	fmt.Scanln(&start, &end)

	feederCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			time.Sleep(300 * time.Millisecond)
			feederCh <- no
		}
		close(feederCh)
	}()

	wg := &sync.WaitGroup{}
	primesCh := make(chan int)
	doneCh := primePrimes(primesCh)
	workerCount := 5
	for id := range workerCount {
		wg.Add(1)
		go func() {
			fmt.Printf("Worker : %d starting...\n", id)
			defer wg.Done()
			defer fmt.Printf("Worker : %d shutting down...\n", id)
			for no := range feederCh {
				fmt.Printf("Worker : %d processing %d...\n", id, no)
				checkPrime(no, primesCh)
			}
		}()
	}
	wg.Wait()
	close(primesCh)
	<-doneCh
}

func primePrimes(primesCh <-chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		for primeNo := range primesCh {
			fmt.Println("Prime No :", primeNo)
		}
		// doneCh <- struct{}{}
		close(doneCh)
	}()
	return doneCh

}

func checkPrime(no int, ch chan<- int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	ch <- no
}
