package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
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

	// also need to declare echange
	err = messageChannel.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		log.Fatal(err)
	}

	q, err := messageChannel.QueueDeclare(
		"",
		true,
		false,
		true, // exclusive // what do mean by it ??
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	// our given queue will do bindings with our exchange.
	err = messageChannel.QueueBind(
		q.Name, // queue name
		"post",     // routing key
		"logs_direct", // exchange
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Now lets consume the message from the given temporary queue
	msgs, err := messageChannel.Consume(q.Name, "", false, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for m := range msgs {
			log.Printf("%s", m.Body)
			m.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	select {}

}
