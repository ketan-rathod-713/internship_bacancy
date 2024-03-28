package app

import (
	"os"
	"rockpaperscissors/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB *mongo.Database
}

func NewApp(db *mongo.Database) *App {
	return &App{
		DB: db,
	}
}

func LoadEnv(prod bool, path string) (*models.Env, error) {
	err := godotenv.Load(path)

	if err != nil {
		return nil, err
	}

	var env models.Env
	if prod {
		env = models.Env{
			PORT:        os.Getenv("PORT"),
			DB_URL:      os.Getenv("DB_URL"),
			SSL_MODE:    os.Getenv("SSL_MODE"),
			DB_NAME:     os.Getenv("DB_NAME"),
			DB_USER:     os.Getenv("DB_USER"),
			DEV_DB_NAME: os.Getenv("DEV_DB_NAME"),
		}
	} else {
		// start development server
		env = models.Env{
			PORT:        os.Getenv("DEV_PORT"),
			DB_URL:      os.Getenv("DEV_DB_URL"),
			SSL_MODE:    os.Getenv("DEV_SSL_MODE"),
			DB_NAME:     os.Getenv("DEV_DB_NAME"),
			DB_USER:     os.Getenv("DEV_DB_USER"),
			DEV_DB_NAME: os.Getenv("DEV_DB_NAME"),
		}
	}

	return &env, nil
}
