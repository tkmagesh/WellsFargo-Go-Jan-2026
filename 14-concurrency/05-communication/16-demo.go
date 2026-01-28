/* using a channel for synchronization (like wait group) */

package main

import (
	"fmt"
	"time"
)

var doneCh chan bool = make(chan bool)

func main() {
	go doSomething()
	fmt.Println("Waiting....")
	<-doneCh
	fmt.Println("Thank you")
}
func doSomething() {
	for range 20 {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	doneCh <- true
}
