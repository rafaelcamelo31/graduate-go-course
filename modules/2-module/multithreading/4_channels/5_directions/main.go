package main

import "fmt"

/*
Go channel direction:
- chan <- This is a receive-only channel
- <- chan This is a send-only channel
*/
func main() {
	ch := make(chan string)
	go read("go direction", ch)

	write(ch)
}

func read(name string, ch chan<- string) {
	ch <- name
}

func write(data <-chan string) {
	fmt.Println(<-data)
}
