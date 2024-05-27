package main

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Data struct {
	EmailTasks []EmailTask `json:"emailTasks"`
}

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

	var data Data = Data{
		EmailTasks: []EmailTask{
			EmailTask{EmailId: "ketanrtd1@gmail.com", Content: "You are ketan rathod"},
			EmailTask{EmailId: "ketanrtd713@gmail.com", Content: "You are ketan 713 ha ha"},
			EmailTask{EmailId: "ketanrtd1@gmail.com", Content: "You are ketan 1 ha ha"},
			EmailTask{EmailId: "ketanrtd1@gmail.com", Content: "You are ketan 1 ha ha"},
		},
	}

	// send emails to given workers and may be email send can be error too.
	for _, d := range data.EmailTasks {
		body, err := json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}
		err = ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}
