package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

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

	todoItem, err := client.TrialTodo(context.Background(), &pb.Noparams{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Trial todo is", todoItem)

	// Call CreateTodo method
	todoItem, err = client.CreateTodo(context.Background(), &pb.CreateTodoRequest{Text: "great"})
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

	// get data from todo stream

	clientStream, err := client.GetTodosStream(context.Background(), &pb.Noparams{})
	if err != nil {
		log.Fatal("error getting client stream for todo items")
	}

	// obviusly error would be nil before starting it
	for {
		data, err := clientStream.Recv()

		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Stream Ended", err)
				break
			} else {
				log.Println("Something wrong happend with client stream", err)
				break
			}
		}

		fmt.Println(data)
	}

	fileStream, err := client.GetFileDownload(context.Background(), &pb.Noparams{})
	if err != nil {
		log.Fatal("error getting file stream", err)
	}

	var setFileName bool = false
	var file *os.File

	for {
		data, err := fileStream.Recv()

		if setFileName == false {
			log.Println("File is created")
			file, err = os.Create(data.Filename)
			setFileName = true
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Stream Ended", err)
				break
			} else {
				log.Println("Something wrong happend with client stream", err)
				break
			}
		}

		fmt.Println("File data has arrived")
		// fmt.Println(data)
		file.Write(data.Chunk)

	}

	defer file.Close()
}
