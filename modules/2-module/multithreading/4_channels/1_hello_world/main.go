package main

import "log"

// Thread 1
func main() {
	c := make(chan string) // Empty channel

	// Thread 2
	go func() {
		c <- "Hello, world!" // Filled channel
	}()

	// Thread 1
	msg := <-c // Empty again
	log.Println(msg)
}
