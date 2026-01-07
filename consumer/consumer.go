package main

import (
	"fmt"

	// RabbitMQ AMQP client
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")

	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Error occured while connection:", err)
		return
	}
	defer conn.Close()

	// Create a channel from connection
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Channel error:", err)
		return
	}
	defer ch.Close()

	// Start consuming messages from queue
	msgs, err := ch.Consume(
		"TestQueue", // queue name
		"",          // consumer tag (auto-generated)
		true,        // auto-ack (message marked done automatically)
		false,       // not exclusive
		false,       // no-local (unused)
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		fmt.Println("Consume error:", err)
		return
	}

	// Channel to block main function forever
	forever := make(chan bool)

	// Goroutine to read messages continuously
	go func() {
		i := 1
		for d := range msgs { // waits for messages from RabbitMQ
			fmt.Printf("Received Message %s : %d\n", string(d.Body), i)
			i++
		}
	}()

	fmt.Println("Connected to Rabbit MQ instance")
	fmt.Println("Waiting for messages...")

	// Block main thread (keeps consumer alive)
	<-forever
}
