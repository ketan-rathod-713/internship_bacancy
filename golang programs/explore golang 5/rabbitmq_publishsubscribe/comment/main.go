package main

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

// The meaning of a binding key depends on the exchange type. The fanout exchanges, which we used previously, simply ignored its value.
// take onnly comment_created service message and process it

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// declare queue to listen to.
	// queue for recieving comments.
	q, err := ch.QueueDeclare("comment_direct", true, false, false, false, nil)

	// bind queue with logs exchange
	err = ch.QueueBind(q.Name, "comment", "logs_direct", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	messages, err := ch.ConsumeWithContext(context.Background(), q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range messages {
		log.Printf("%s", msg.Body)
		// msg.Ack(true)
	}

	select {}
}
