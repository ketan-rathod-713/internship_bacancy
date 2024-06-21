package api

import (
	"encoding/json"
	"net/http"
	"rockpaperscissors/api/user"
	"rockpaperscissors/app"

	"github.com/gorilla/mux"
)

// Separate all the api's here.
type api struct {
	App *app.App
}

func NewApi(app *app.App) *api {
	return &api{
		App: app,
	}
}

func (a *api) InitializeRoutes(app *mux.Router) {
	userApi := user.New(a.App.DB)

	a.ConnectToSocket(app)

	// http middleware to fiber handler
	// app.Use(adaptor.HTTPMiddleware(middlewares.LogMiddleware))

	// api status
	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Status: 200, Message: "Api working fine", Data: nil})
	})

	apiRoute := app.PathPrefix("/api")

	userRoute := apiRoute.PathPrefix("/user")
	userApi.Routes(userRoute.Subrouter())

}
