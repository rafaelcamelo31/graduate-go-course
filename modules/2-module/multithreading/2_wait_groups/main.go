package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := range 10 {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1
func main() {
	wg := sync.WaitGroup{}
	wg.Add(25)
	// Thread 2
	go task("A", &wg)

	// Thread 3
	go task("B", &wg)

	// Thread 4
	go func() {
		for i := range 5 {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			wg.Done()
		}
	}()
	wg.Wait()
}
