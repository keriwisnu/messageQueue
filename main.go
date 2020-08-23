package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println(" - This is Publisher -")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connect to rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MessageQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"MessageQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("This message from publisher for consumer"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Message Published to queue")

}
