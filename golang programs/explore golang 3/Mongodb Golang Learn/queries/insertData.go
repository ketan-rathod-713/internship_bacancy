package queries

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	CustomerId       string `csv:"Customer Id"`
	FirstName        string `csv:"First Name"`
	LastName         string `csv:"Last Name"`
	Company          string `csv:"Company"`
	City             string `csv:"City"`
	Country          string `csv:"Country"`
	Email            string `csv:"Email"`
	SubscriptionDate string `csv:"Subscription Date"`
}

func InsertRequiredData(client *mongo.Client) {
	customer := client.Database("mongodblearn").Collection("customer")

	file, err := os.Open("data/customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)

	for {
		var record []string
		record, err = csvReader.Read()

		if errors.Is(err, io.EOF) {
			log.Println("Completed reading file")
			break
		}

		// push record to mongodb for further processing
		var customerStruct Customer = Customer{
			CustomerId:       record[1],
			FirstName:        record[2],
			LastName:         record[3],
			Company:          record[4],
			City:             record[5],
			Country:          record[6],
			Email:            record[9],
			SubscriptionDate: record[10],
		}

		result, err := customer.InsertOne(context.TODO(), customerStruct)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(result.InsertedID)
	}
}
