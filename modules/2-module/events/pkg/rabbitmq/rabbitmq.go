package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, output chan amqp.Delivery, queueName string) error {
	msgs, err := ch.Consume(queueName, "goConsumer", false, false, false, false, nil)
	if err != nil {
		return err
	}
	for msg := range msgs {
		output <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, exchangeName string) error {
	err := ch.Publish(exchangeName, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		return nil
	}
	return nil
}
