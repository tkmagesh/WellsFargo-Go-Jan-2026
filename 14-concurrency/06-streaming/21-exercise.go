package main

import (
	"fmt"
	"sync"
)

/*
Refactor the below to follow "Share memory by communicating"
*/

var primes []int
var mutex sync.Mutex

func main() {
	var start, end int
	fmt.Println("Enter the start and end")
	fmt.Scanln(&start, &end)
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkPrime(no, wg)
	}
	wg.Wait()
	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}
}

func checkPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primes = append(primes, no)
	}
	mutex.Unlock()
}
