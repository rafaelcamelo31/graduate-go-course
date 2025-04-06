package main

func main() {
	// Create a channel
	ch := make(chan int)

	go publish(ch)

	reader(ch)
}

func reader(ch chan int) {
	for i := range ch {
		println("Received", i)
	}
}

func publish(ch chan int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}
