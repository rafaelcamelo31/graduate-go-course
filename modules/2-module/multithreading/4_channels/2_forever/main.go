package main

func main() {
	// Empty channel
	forever := make(chan bool)

	go func() {
		for i := range 10 {
			println(i)
		}
		// Filled channel
		forever <- true
	}()

	// Channel is filled, so the process ends without deadlock
	<-forever
}
