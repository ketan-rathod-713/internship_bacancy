package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
	_ "github.com/lib/pq"                                // Import the PostgreSQL driver
)

func main() {
	var m bool
	var migrationType string
	var n int

	flag.BoolVar(&m, "migrate", false, "for migrating or not")
	flag.StringVar(&migrationType, "type", "up", "for up or down migration")
	flag.IntVar(&n, "n", 1, "for up migration")

	flag.Parse()

	if m {
		log.Println("Migrate database")

		db, err := sql.Open("postgres", "postgresql://bacancy:admin@localhost/bacancy?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatal(err)
		}

		m, err := migrate.NewWithDatabaseInstance("file://db/migration/", "postgres", driver)
		if err != nil {
			log.Fatal(err)
		}

		if migrationType == "up" {
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		} else if migrationType == "down" {
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		} else if migrationType == "steps" {
			if err := m.Steps(n); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		} else {
			log.Fatal("Invalid Migration Type use up or down only	")
		}

		log.Println("Migration successful")

		v, dirty, _ := m.Version()
		log.Println("current db version", v, dirty)

	} else {
		log.Println("Normal Code Here")
	}
}
