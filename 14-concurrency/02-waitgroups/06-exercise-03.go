package main

import (
	"fmt"
	"sync"
)

/*
Accept "start" and "end" from the user (numbers)
Print all the prime numbers between "start" and "end" (including start and end)
*/

type Concur struct {
	sync.WaitGroup
}

func (cc *Concur) Run(f func()) {
	cc.Add(1)
	go func() {
		defer cc.Done()
		f()
	}()
}

func main() {
	var start, end int
	fmt.Println("Enter the start and end")
	fmt.Scanln(&start, &end)
	cc := &Concur{}
	for no := start; no <= end; no++ {
		cc.Run(func() {
			checkPrime(no)
		})
	}
	cc.Wait()

}

func checkPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
