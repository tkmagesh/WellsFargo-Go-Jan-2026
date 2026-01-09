package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("Spinninng %d goroutines.. Hit ENTER to start...\n", count)
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	for id := range count {
		wg.Add(1)
		go fn(wg, id)
	}

	wg.Wait() // block the execution until the wg counter becomes 0 (default = 0)
	fmt.Println("Done!")
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)

}
