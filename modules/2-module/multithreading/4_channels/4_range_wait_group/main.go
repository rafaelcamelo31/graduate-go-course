package main

import "sync"

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)

	go reader(ch, &wg)

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		println("Received", i)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := range 10 {
		ch <- i
	}
}
