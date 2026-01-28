package main

import (
	"fmt"
	"sync"
)

/*
Accept "start" and "end" from the user (numbers)
Print all the prime numbers between "start" and "end" (including start and end)
*/

func main() {
	var start, end int
	fmt.Println("Enter the start and end")
	fmt.Scanln(&start, &end)
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Go(func() {
			checkPrime(no)
		})
	}
	wg.Wait()

}

func checkPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
