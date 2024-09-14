package main

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EmailTask struct {
	EmailId string `json:"emailId"`
	Content string `json:"content"`
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// AMQP queue
	q, err := ch.QueueDeclare("email_task_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	messagesChannel, err := ch.ConsumeWithContext(context.Background(), q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for m := range messagesChannel {

			var emailTask EmailTask
			err = json.Unmarshal(m.Body, &emailTask)
			if err != nil {
				continue // handle next request.
			}
			log.Println("before processing",emailTask)
			// for email task send email to each one.
			if emailTask.EmailId == "ketanrtd1@gmail.com" {
				continue // handle next request
			}

			log.Println(emailTask)

			m.Ack(true)
		}
	}()

	select {}
}
