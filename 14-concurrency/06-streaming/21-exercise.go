package main

import (
	"fmt"
	"sync"
)

/*
Refactor the below to follow "Share memory by communicating"
*/

func main() {
	var start, end int
	fmt.Println("Enter the start and end")
	fmt.Scanln(&start, &end)
	wg := &sync.WaitGroup{}
	primesCh := make(chan int)
	doneCh := primePrimes(primesCh)
	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkPrime(no, wg, primesCh)
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

func checkPrime(no int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	ch <- no
}
