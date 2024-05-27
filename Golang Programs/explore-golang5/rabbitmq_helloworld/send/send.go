package main

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Error Connecting To Rabbit Mq")
	defer conn.Close()

	// The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication and so on for us. Next we create a channel, which is where most of the API for getting things done resides:

	ch, err := conn.Channel()
	failOnError(err, "Error creating channel in amqp")
	defer ch.Close()

	// To send, we must declare a queue for us to send to; then we can publish a message to the queue:

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "error declaring queue in rabbit mq")

	body := "Hello World!"
	err = ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
