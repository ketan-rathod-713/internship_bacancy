package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	URL := `postgresql://bacancy:admin@localhost/bacancy`

	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Println(err)
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
	}

	defer db.Close() // Pushed in stack

	log.Println("Connected!")

	// transaction
}

func InsertBook(db *sql.DB, bookName string, bookQuantity int) {
	query := "INSERT INTO transaction.book(name, quantity) VALUES($1, $2);"
	db.Begin()
	db.Exec(query, bookName, bookQuantity)
	db.Close()
}

func IssueBook(db *sql.DB, bookId int) bool{
	
}
