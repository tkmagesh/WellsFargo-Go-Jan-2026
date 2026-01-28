package main

import (
	"fmt"
	"sync"
)

/*
Refactor the below in such a way that the identified prime numbers are printed in the "main" function (not in the checkprime function)


*/

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
}

func checkPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
