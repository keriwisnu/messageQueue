package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println(" - This is Consumer -")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	mssg, err := ch.Consume(
		"MessageQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	devour := make(chan bool)
	go func() {
		for x := range mssg {
			fmt.Printf("Received Message : %s\n", x.Body)
		}
	}()
	fmt.Println("Devour message from queue")
	fmt.Println(" - waiting the message -")
	<-devour
}
