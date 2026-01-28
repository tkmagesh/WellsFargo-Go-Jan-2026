package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateData()
	fmt.Println("data :", <-ch)
	fmt.Println("data :", <-ch)
	fmt.Println("data :", <-ch)
	fmt.Println("data :", <-ch)
	fmt.Println("data :", <-ch)
	fmt.Println("Done")
}

func generateData() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 10
		time.Sleep(2 * time.Second)
		ch <- 20
		time.Sleep(2 * time.Second)
		ch <- 30
		time.Sleep(2 * time.Second)
		ch <- 40
		time.Sleep(2 * time.Second)
		ch <- 50
	}()
	return ch
}
