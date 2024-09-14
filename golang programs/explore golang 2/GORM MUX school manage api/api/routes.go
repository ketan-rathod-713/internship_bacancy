package api

import (
	"net/http"
	parentapi "schoolApi/api/parentApi"
	sportapi "schoolApi/api/sportApi"
	studentapi "schoolApi/api/studentApi"
	"schoolApi/app"

	"github.com/gorilla/mux"
)

// What are different routes

// for given school

// /student GET get all students
// /student POST post one student
// PUT
// Patch etc.

type Api struct {
	App *app.App
}

func NewApi(app *app.App) *Api {
	return &Api{App: app}
}

func (a *Api) InitialiseRoutes(router *mux.Router) {
	router.HandleFunc("/", homeHandler)

	studentRouter := router.PathPrefix("/student").Subrouter()
	studentapi.Routes(studentRouter, a.App.DB)

	sportRouter := router.PathPrefix("/sport").Subrouter()
	sportapi.Routes(sportRouter, a.App.DB)

	parentRouter := router.PathPrefix("/parent").Subrouter()
	parentapi.Routes(parentRouter, a.App.DB)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hellow ha hah"))
}
