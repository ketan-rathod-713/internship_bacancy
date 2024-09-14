package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

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

func (s *todoserver) TrialTodo(ctx context.Context, req *pb.Noparams) (*pb.TodoItem, error) {

	return &pb.TodoItem{
		Title:     "Trial todo",
		TrialTodo: true,
		Completed: false,
		Id:        1,
	}, nil
}

func (s *todoserver) GetTodos(ctx context.Context, _ *pb.Noparams) (*pb.TodoItems, error) {
	return &pb.TodoItems{TodoItems: todos}, nil
}

// here we have slight different syntax
// first it takes the arguments and then it takes the server stream which we are going to utilize to send the data to client
func (s *todoserver) GetTodosStream(_ *pb.Noparams, stream pb.Todo_GetTodosStreamServer) error {

	for i, v := range todos {
		stream.Send(v)
		log.Println("Todo Index Processed", i)
		time.Sleep(time.Second)
	}

	log.Println("Closing Todo Stream")
	// How to close the stream ? or do it automatically gets closed

	return nil
}
func (s *todoserver) GetFileDownload(_ *pb.Noparams, stream pb.Todo_GetFileDownloadServer) error {

	// read from file and send content in byte to client
	file, err := os.Open("../public/file1.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	var buffer []byte = make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("Error reading content from file")
			break
		}

		// send data to client
		stream.Send(&pb.TypeFileDownload{
			Filename: file.Name(),
			Chunk:    buffer[:n],
		})
	}
	// How to close the stream ? or do it automatically gets closed

	fmt.Println("Closing file reading stream")
	return nil
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
