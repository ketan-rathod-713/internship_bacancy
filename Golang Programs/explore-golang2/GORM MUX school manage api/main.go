package main

import (
	"fmt"
	"log"
	"net/http"
	"schoolApi/api"
	"schoolApi/app"

	"github.com/gorilla/mux"
)

func main() {
	app, err := app.NewApp()

	if err != nil {
		log.Println("AN error occured")
	}

	a := api.NewApi(app)

	router := mux.NewRouter()

	a.InitialiseRoutes(router)

	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.PORT), router)
}
