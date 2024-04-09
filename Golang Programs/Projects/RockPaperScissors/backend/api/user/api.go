package user

import (
	"rockpaperscissors/app/gameservice"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	DB      *mongo.Database
	Service gameservice.Service
}

func New(db *mongo.Database) *api {
	return &api{DB: db, Service: gameservice.New(db)}
}

// define user related routes
func (a *api) Routes(parentRouter *mux.Router) {
	parentRouter.HandleFunc("/signin", a.SignInUser).Methods("POST") // get userId and Password
	parentRouter.HandleFunc("/", a.GetUsers).Methods("GET")
	parentRouter.HandleFunc("/", a.PostUser).Methods("POST")

	// auth related functionality

	// generate auth token and give it in json and also set cookie of it.
	// parentRouter.Post("/login", a.Login)
	// sign up is already done by create user api.
}
