package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var ctx context.Context

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

	ctx = context.Background() // never dieas and basic context not nil empty context
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO sql.emp(name, age) VALUES('Ketan', 89);")
	if err != nil {
		tx.Rollback()
		fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
		return
	}

	fmt.Println("First Query Ran")

	// The next query is handled similarly
	_, err = tx.ExecContext(ctx, "UPDATE sql.emp SET age = 0 WHERE age = 23 somethinglfdgkqsjndjgnaq")
	if err != nil {
		tx.Rollback()
		fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("....Transaction committed\n")
	}

	// To start transaction without specifying context
	tx1, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	// tx1.Exec()
	// tx1.Exec()

	tx1.Commit() // then commit or rollbac based on queries results

	// You can define a prepared statement for repeated use. This can help your code run a bit faster by avoiding the overhead of re-creating the statement each time your code performs the database operation.

	// func (db *DB) Prepare(query string) (*Stmt, error)

	stmt, err := db.Prepare("INSERT INTO sql.emp(name, age) VALUES( $1, $2 )")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("ketan", 29)
	if err != nil {
		log.Println(err)
	}

	// Query Row
	// TODO: Work on isolation levels and so on.
}
