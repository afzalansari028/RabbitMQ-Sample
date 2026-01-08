package main

import (
	"fmt"

	// RabbitMQ AMQP client for Go
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Rabbit MQ!!!!!!!!")

	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	// Create a channel from the connection
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Channel error:", err)
		return
	}
	defer ch.Close()

	// Declare queue (creates if not exists)
	q, err := ch.QueueDeclare(
		"TestQueue", // queue name
		false,       // durable - restart lost
		false,       // auto delete
		false,       // exclusive 
		false,       // no wait
		nil,         // args
	)
	if err != nil {
		fmt.Println("Queue error:", err)
		return
	}
	fmt.Println("Queue:", q.Name)

	// Publish message to queue (default exchange)
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println("Message sent")
}
