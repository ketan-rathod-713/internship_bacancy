package main

import (
	"context"
	"log"
	"mondodblearn/queries"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal("Error connecting mongodb")
	}

	// customer := client.Database("mongodblearn").Collection("customer")
	// queries.InsertRequiredData(client)

	queries.GroupByState(client)
}
