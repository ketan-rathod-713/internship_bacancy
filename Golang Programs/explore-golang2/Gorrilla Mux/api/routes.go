package api

import (
	"net/http"
	"postgres-crud/app"

	"github.com/gorilla/mux"
)

type API struct {
	App *app.App
}

// Create new APi
func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	return api, nil
}

// initialsie some routes on given router on given api
func (a *API) InitializeRoutes(router *mux.Router) {
	// create different services api and handle them with routes.
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world ha ha"))
	}).Methods("GET")

	// TODO: Move this method to booksApi. 
	// TODO: Also add customerApi and move its logic in customerAPI
}
