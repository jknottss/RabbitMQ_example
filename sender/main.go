package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channalRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channalRabbitMQ.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channalRabbitMQ.QueueDeclare(
		"QueueService1", //queue name
		true,            // durable
		false,           //auto delete
		false,           //exclusive
		false,           //no wait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	//Create a new Fiber instance
	app := fiber.New()

	// Add middleware
	app.Use(
		logger.New(), //add a simple logger
	)

	// Add route
	app.Get("/send", func(c *fiber.Ctx) error {
		// Create message to publish
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("msg")),
		}

		// Attempt to publish a message to the queue
		if err := channalRabbitMQ.Publish(
			"",              //exchange
			"QueueService1", //queue name
			false,           //mandatory
			false,           // immediate
			message,         //massage to publish
		); err != nil {
			return err
		}
		return nil
	})

	// Start Fiber API server
	log.Fatal(app.Listen(":3000"))
}
