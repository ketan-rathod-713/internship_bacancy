package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcproject/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("did not connect %v", err)
	}

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NameList{
	// 	Name: []string{"Akhil", "aman", "ketan"},
	// }

	callSayHello(client)
}
