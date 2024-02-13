package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitialiseDB() (*sql.DB, error) {
	URL := "postgres://bacancy:admin@localhost/bacancy?sslmode=disable"

	db, err := sql.Open("postgres", URL)

	if err != nil {
		return nil, err
	}

	if db.Ping() != nil {
		return nil, err
	}

	return db, nil
}
