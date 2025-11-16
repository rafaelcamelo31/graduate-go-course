package main

import (
	"fmt"
	"time"
)

func main() {
	// channel with buffer
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		fmt.Println("Sent 1")
		ch <- 2
		fmt.Println("Sent 2")
		// This will be blocked until there is a space in buffer
		ch <- 3
		fmt.Println("Sent 3")
		close(ch)
	}()

	time.Sleep(time.Second)

	for v := range ch {
		fmt.Println("Received", v)
	}

	time.Sleep(2 * time.Second)
}
