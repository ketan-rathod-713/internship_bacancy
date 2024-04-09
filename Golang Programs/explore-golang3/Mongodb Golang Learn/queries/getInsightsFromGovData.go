package queries

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
? get state wise data
*/

func GroupByState(client *mongo.Client) {
	govdata := client.Database("mongodblearn").Collection("govdata")
	filter := bson.M{
		"state_code": 24,
	}

	cursor, err := govdata.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var result []map[string]interface{} = make([]map[string]interface{}, 0)
	cursor.All(context.TODO(), &result)

	fmt.Println(result[0])
}

func PipelineConcept1(client *mongo.Client) {
	govdata := client.Database("mongodblearn").Collection("govdata")

	pipeline := []bson.M{
		{
			"$addFields": bson.M{
				"homework": "$state_code",
			},
		},
	}

	cursor, err := govdata.Aggregate(context.TODO(), pipeline)

	if err != nil {
		log.Fatal(err)
	}

	var result []map[string]interface{} = make([]map[string]interface{}, 0)
	cursor.All(context.TODO(), &result)

	fmt.Println(result[0])
}
