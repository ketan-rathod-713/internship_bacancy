package main

import (
	"facebook/api"
	"facebook/database"
	"facebook/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	configs, err := database.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB(configs)
	if err != nil {
		log.Fatal(err)
	}

	resolver := &graph.Resolver{
		API: api.NewApi(db.Database("graphqlfriends"), configs),
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost%s/ for GraphQL playground", configs.PORT)
	log.Fatal(http.ListenAndServe(configs.PORT, nil))
}
