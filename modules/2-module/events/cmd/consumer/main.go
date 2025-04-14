package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "firstQueue")

	fmt.Println("Running server...")
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
