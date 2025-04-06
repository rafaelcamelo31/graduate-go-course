package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64

	// RabbitMQ
	go func() {
		for {
			time.Sleep(2 * time.Second)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg

		}
	}()

	// Kafka
	go func() {
		for {
			time.Sleep(2 * time.Second)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg.id, msg.Msg)

		case msg := <-c2:
			fmt.Printf("Received from Kafka: ID: %d - %s\n", msg.id, msg.Msg)

		case <-time.After(3 * time.Second):
			println("Timeout")

			// default:
			// 	// This case will be executed if no other case is ready
			// 	println("No messages received")
		}
	}
}
