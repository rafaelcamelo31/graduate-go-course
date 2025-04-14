package main

import (
	"github.com/rafaelcamelo31/graduate-go-course/2-module/events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, RabbitMQ.", "amq.direct")
}
