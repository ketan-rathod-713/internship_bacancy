package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpcproject/proto/todo"
)

type todoserver struct {
	pb.UnimplementedTodoServer
}

var todos []*pb.TodoItem

// wow it was so easy to implement grpc in golang
func (s *todoserver) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.TodoItem, error) {
	todo := &pb.TodoItem{
		Id:        1,
		Title:     req.GetText(),
		Completed: false,
	}
	todos = append(todos, todo)
	return todo, nil
}

func (s *todoserver) GetTodos(ctx context.Context, _ *pb.Noparams) (*pb.TodoItems, error) {
	return &pb.TodoItems{TodoItems: todos}, nil
}

// CreateTodo(context.Context, *CreateTodoRequest) (*TodoItem, error)
// GetTodos(context.Context, *Noparams) (*TodoItems, error)

func main() {
	fmt.Println("Server is starting")

	server := grpc.NewServer()
	s := todoserver{}

	// listen on port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterTodoServer(server, &s)

	// Start the server
	log.Println("Server started listening on port 8080...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
