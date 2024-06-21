package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/lib/pq"
)

type todo struct {
	id        sql.NullInt32
	title     sql.NullString
	completed sql.NullBool
	createdAt sql.NullString
}

// connect to database

func connect() *sql.DB {
	var err error
	db, err := sql.Open("postgres", "postgres://bacancy:admin@localhost/bacancy?sslmode=disable")
	checkError(err)

	checkError(db.Ping())

	// fmt.Println("Connected To Database")

	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	db := connect()
	// createTodoTable()

	if len(os.Args) >= 2 && os.Args[1] == "help" {
		fmt.Println("Functions Available :")
		fmt.Println("1. list : To view all the todos")
		fmt.Println("2. create <name of todo> : To create todo with given name")
	}

	if len(os.Args) >= 2 && os.Args[1] == "list" {
		showTodos(getTodos(db))
	}

	if len(os.Args) >= 3 && os.Args[1] == "create" { // give 2 arguments create titleName
		createTodo(db, os.Args[2])
	}

	if len(os.Args) >= 3 && os.Args[1] == "completed" { // give 2 arguments create titleName
		id , err:= strconv.Atoi(os.Args[2])
		checkError(err)
		completedTodo(db, id)
	}
}

func completedTodo(db *sql.DB, id int) {
	query := `UPDATE todo_cli.todos
	SET completed = true
	WHERE id = $1;`
	_, err := db.Exec(query, id)
	checkError(err)
}

func createTodo(db *sql.DB, title string) {
	query := `INSERT INTO todo_cli.todos(title, completed, createdat) VALUES($1, false, NOW());`
	_, err := db.Exec(query, title)
	checkError(err)
}

func createTodoTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS todo_cli.todos (id SERIAL, title VARCHAR(255), completed boolean, createdat date);`
	_, err := db.Exec(query)
	checkError(err)
}

func getTodos(db *sql.DB) []todo {
	var todos []todo
	query := `SELECT id, title, completed, createdat FROM todo_cli.todos;`
	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var todo todo = todo{}
		rows.Scan(&todo.id, &todo.title, &todo.completed, &todo.createdAt)
		todos = append(todos, todo)
	}

	return todos
}

func showTodos(todos []todo) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Completed", "Created At"})
	// t.AppendRows([]table.Row{
	// 	{1, "Arya", "Stark", 3000},
	// 	{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	// })
	// t.AppendFooter(table.Row{"", "", "Total", 10000})

	// fmt.Println("\n\tTODOs\n")
	// fmt.Printf("%v\t\t%v\t%v\t%v\n", "N", "Title", "Completed", "Created At")
	for _, todo := range todos {
		t.AppendRow([]interface{}{todo.id.Int32, todo.title.String, todo.completed.Bool, todo.createdAt.String})
		// fmt.Printf("%v\t\t%v\t%v\t\t%v\n", todo.id.Int32, todo.title.String, todo.completed.Bool, todo.createdAt.String)
	}
	t.Render()
}
