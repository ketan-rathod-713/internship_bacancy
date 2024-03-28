package main

import (
	"context"
	pb "grpcproject/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatal("could not greet", err)
	}

	log.Printf("%s", res.Message)
}
