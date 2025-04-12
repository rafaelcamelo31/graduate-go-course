package main

import "fmt"

func main() {
	events := []string{"test1", "test2", "test3", "test4"}
	events = append(events[:0], events[1:]...)
	fmt.Println(events)
}
