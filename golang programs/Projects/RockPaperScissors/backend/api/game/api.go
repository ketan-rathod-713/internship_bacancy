package game

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	DB *mongo.Database
}

func New(db *mongo.Database) *api {
	return &api{DB: db}
}

// define user related routes
func (a *api) Routes(parentRouter fiber.Router) {

	// here id is the userId
	parentRouter.Get("/:id", a.GetGamesPlayedByUser)
}
