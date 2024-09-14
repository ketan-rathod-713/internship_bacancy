package api

import (
	"graphql/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	DB     *mongo.Database
	Config *database.Config
}

func NewApi(db *mongo.Database, config *database.Config) *Api {
	return &Api{DB: db, Config: config}
}
