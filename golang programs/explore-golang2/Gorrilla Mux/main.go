package main

import (
	"log"
	"net/http"
	"postgres-crud/api"
	"postgres-crud/app"

	"github.com/gorilla/mux"
)

func main() {
	// Create App Which Will Initialise Database and populate pointer of db inside it DB
	app := app.NewApp()
	log.Println(app)

	// Initialize All APIs
	api, err := api.New(app)
	if err != nil {
		log.Println("API INITIALIZATION ERROR")
	}
	log.Println(api.App)

	// Now initialise routes on given router
	var router *mux.Router = mux.NewRouter()
	api.InitializeRoutes(router)

	// Now listen and serve
	http.ListenAndServe(":8080", router)
}

