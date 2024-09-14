package main

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	// channel opens new concurrent channel to amqp server and we can process bulk of amqp messages to go.
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("worker", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(q.Consumers, q.Messages, q.Name)

	mssagesChannel, err := ch.ConsumeWithContext(context.Background(), q.Name, "", false, false, false, false, nil)

	var forever chan struct{}

	go func(){
		for d := range mssagesChannel {
			log.Printf("Recieved message %s", d.Body)
			d.Ack(true)
		}
	}()


	<-forever
}
