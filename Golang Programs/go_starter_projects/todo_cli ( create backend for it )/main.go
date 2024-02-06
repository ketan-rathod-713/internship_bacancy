package main

import (
	"fmt"
	"os"
	"time"
)

type todo struct {
	title     string
	completed bool
	createdAt string
}

func main() {
	todos := []todo{{"First todo", false, "12-01-2024"}}

	if len(os.Args) >= 2 && os.Args[1] == "show" {
		showTodos(todos)
	}

	if len(os.Args) >= 3 && os.Args[1] == "create" { // give 2 arguments create titleName
		title := os.Args[2]
		todos = append(todos, todo{title, false, time.Now().String()})
	}
}

func showTodos(todos []todo) {
	fmt.Printf("%v \t %v \t %v \t %v\n", "Index", "Completed", "Title", "Created At")
	for index, todo := range todos {
		fmt.Printf("%v \t %v \t %v \t %v\n", index, todo.completed, todo.title, todo.createdAt)
	}
}
