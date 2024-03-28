package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpcproject/proto/todo" // Import your generated protobuf package
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a Todo client
	client := pb.NewTodoClient(conn)

	// Call CreateTodo method
	todoItem, err := client.CreateTodo(context.Background(), &pb.CreateTodoRequest{Text: "great"})
	if err != nil {
		log.Fatalf("CreateTodo failed: %v", err)
	}
	log.Printf("Created Todo item: %v", todoItem)

	// Call ReadTodos method
	todoList, err := client.GetTodos(context.Background(), &pb.Noparams{})
	if err != nil {
		log.Fatalf("ReadTodos failed: %v", err)
	}
	log.Println("Todo List:")
	for _, todo := range todoList.TodoItems {
		log.Printf("ID: %d, Text: %s", todo.Id, todo.Title)
	}
}
