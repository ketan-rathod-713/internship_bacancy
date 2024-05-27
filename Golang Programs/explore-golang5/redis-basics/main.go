package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	log.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(pong)

	// client.Set(context.TODO(), "data", "good", 10*time.Second)
	val, err := client.Get(context.TODO(), "data").Result()

	log.Println(val, err)

}
