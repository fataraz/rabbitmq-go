package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	// connect to rabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		fmt.Println("err connect", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to rabbitMQ")

	// connect to rabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("err channel", err)
	}
	defer ch.Close()

	// declare queue
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("err queue", err)
	}

	// send data to rabbitMQ with struct
	type Person struct {
		Name     string
		Age      int
		Email    string
		Password string
	}

	Andi := Person{
		Name:     "Andi",
		Age:      10,
		Email:    "andi@gmail.com",
		Password: "123456",
	}
	DataJson, _ := json.Marshal(Andi)

	// publish queue to rabbitMQ
	if err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        DataJson,
		}); err != nil {
		fmt.Println("err publish", err)
	}

	fmt.Println("Successfully publis message to Queue")
}
