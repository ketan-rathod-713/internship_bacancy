package app

import (
	"database/sql"
	"log"
	"postgres-crud/database"
)

// Can be used for app level data handling // mainly for database
type App struct {
	DB *sql.DB
}

func NewApp() *App { // ohkk declared and return bhi yahi ho jaega ?? TODO:
	app := &App{} // Make this before assigning any value using .DB
	var err error
	app.DB, err = database.InitialiseDB()

	if err != nil {
		log.Println("Error Connecting Database")
	}

	return app
}
