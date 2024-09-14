package main

import (
	"graphql/api"
	"graphql/database"
	"graphql/graph"
	"graphql/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	config, err := database.LoadEnv()
	if err != nil {
		log.Fatal("error loading env variables")
	}

	mongoClient, err := database.ConnectDB(config)
	if err != nil {
		log.Fatal("error connecting database")
	}

	// pass this valuable information inside resolver struct for further use
	resolvers := &graph.Resolver{
		// DB: mongoClient.Database(config.DATABASE),
		Api: api.NewApi(mongoClient.Database(config.DATABASE), config),
	}

	// graph folder resolver is used here. so that all type definitions and query and mutations can be tied to it.
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	// Open graphql playground here
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// For quering purposes
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.PORT)
	log.Fatal(http.ListenAndServe(config.PORT, nil))
}
