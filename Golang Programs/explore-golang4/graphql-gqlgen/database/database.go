package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	PORT     string
	HOST     string
	DB_URL   string
	DATABASE string
}

func ConnectDB(config *Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DB_URL))

	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected To Mongodb")
	return mongoClient, nil
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &Config{
		PORT:     os.Getenv("PORT"),
		DB_URL:   os.Getenv("DB_URL"),
		HOST:     os.Getenv("HOST"),
		DATABASE: os.Getenv("DATABASE"),
	}, nil
}
