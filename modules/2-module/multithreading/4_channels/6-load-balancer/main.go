package main

import "time"

func main() {
	data := make(chan int)
	workerQuantity := 100

	// Initialize range of workers
	for i := range workerQuantity {
		go worker(i, data)
	}

	//
	for i := range 1000 {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		println("Worker", workerId, "processing", x)
		time.Sleep(time.Second)
	}
}
