package main

import (
	"flag"
	"fmt"
	"sync"
)

/*
Refactor the below to follow "Share memory by communicating"
*/

func main() {
	var start, end int
	var workerCount int
	flag.IntVar(&workerCount, "workers", 1, "Number of workers to employ")
	flag.Parse()
	fmt.Printf("Employing %d workers\n", workerCount)
	fmt.Println("Enter the start and end")
	fmt.Scanln(&start, &end)
	feederCh := getFeeder(start, end)
	primesCh := startWorkers(workerCount, feederCh)
	doneCh := primePrimes(primesCh)
	<-doneCh
}

func startWorkers(workerCount int, feederCh <-chan int) <-chan int {
	primesCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for id := range workerCount {
			wg.Add(1)
			go func() {
				fmt.Printf("Worker : %d starting...\n", id)
				defer wg.Done()
				defer fmt.Printf("Worker : %d shutting down...\n", id)
				for no := range feederCh {
					if isPrime(no) {
						primesCh <- no
					}
				}
			}()
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func getFeeder(start, end int) <-chan int {
	feederCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			// time.Sleep(300 * time.Millisecond)
			feederCh <- no
		}
		close(feederCh)
	}()
	return feederCh
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

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
