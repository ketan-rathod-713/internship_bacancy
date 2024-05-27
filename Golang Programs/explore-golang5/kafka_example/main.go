package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.7.103:9092",
		"group.id":          "anything",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		log.Println(err)
	}
	err = c.SubscribeTopics([]string{"online"}, nil)
	if err != nil {
		log.Println(err)
	}
	var run bool = true
	for run {
		e := c.Poll(1000)
		switch e.(type) {
		case *kafka.Message:
			log.Println(string(e.(*kafka.Message).Value))
		case kafka.Error:
			log.Println(e.(kafka.Error).Error())
			run = false
			break
		default:
			log.Println("i don't know")
			break
		}
	}
	c.Close()
}
