package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type Event struct {
	Name      string
	TimeStamp string
}

func main() {
	eventName := os.Args[1]
	recieverService := os.Args[2]

	fmt.Println(eventName)

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	messageChannel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer messageChannel.Close()

	// declare queue before using it. // Producer can only sends message to a exchange only // Exchange Work : on one side it recieves the message from the producer and on other side it sends it to queues.
	// 4 exchange types available //  fanout, topic, direct, headers // declare echange before using it.

	// err = messageChannel.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	// fanout doesn't care about routing key. but direct does.

	err = messageChannel.ExchangeDeclare("logs_direct", "direct", true, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	// fanout : broadcasts message to every queue

	// bind queues // one comment, one post and other queue for seeing logs 	// earlier published message on default exchange, now lets declare it on our named exchange.

	// body := "good"
	event := Event{
		Name:      eventName,
		TimeStamp: time.Now().String(),
	}
	body, err := json.Marshal(&event)
	if err != nil {
		log.Fatal(err)
	}

	err = messageChannel.PublishWithContext(context.TODO(), "logs_direct", recieverService, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	},
	)

	if err != nil {
		log.Fatal(
			err,
		)
	}

	// Giving a queue a name is important when you want to share the queue between producers and consumers.

	// But that is not the case for our logger.
	// But that's not the case for our logger. We want to hear about all log messages, not just a subset of them. We're also interested only in currently flowing messages not in the old ones. To solve that we need two things.

	// Firstly, whenever we connect to Rabbit we need a fresh, empty queue. To do this we could create a queue with a random name, or, even better - let the server choose a random queue name for us.

}

// Example for given project.
// I will send query to subscribe that post_created then it will handle it. and also another query like coment created  for comment service.
