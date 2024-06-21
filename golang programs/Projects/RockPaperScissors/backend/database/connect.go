package database

import (
	"context"
	"rockpaperscissors/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitialiseDatabase(env *models.Env) (*mongo.Client, error) {
	// connStr := fmt.Sprintf("user=%v dbname=%v sslmode=%v", env.DB_USER, env.DB_NAME, env.SSL_MODE)
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	return nil, err
	// }

	// err = db.Ping()

	// if err != nil {
	// 	return nil, err
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.DB_URL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
