package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	// Define RabbitMQ server URL
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to QueueService1 for getting messages
	messages, err := channelRabbitMQ.Consume(
		"QueueService1", //queue name
		"",              // consumer
		true,            // auto-ack
		false,           //exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	go func() {
		for message := range messages {
			// For example, show received message in a console
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	// Make a channel to receive massages into infinite loop
	<-make(chan bool)
}
