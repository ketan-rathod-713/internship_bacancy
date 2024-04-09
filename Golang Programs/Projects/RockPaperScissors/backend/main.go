package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"rockpaperscissors/api"
	"rockpaperscissors/app"
	"rockpaperscissors/database"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	var prod bool

	flag.BoolVar(&prod, "startprodserver", false, "By default start development server, if specified this then start production server with its configuration")
	// dev and prod
	flag.Parse()

	env, err := app.LoadEnv(prod, "./.env")
	if err != nil {
		log.Fatal("unable to load env file")
	}

	client, err := database.InitialiseDatabase(env)

	App := app.NewApp(client.Database(env.DB_NAME))
	Api := api.NewApi(App)

	// Now create api and then initialise routes
	// REMOVE Fiber
	// app := fiber.New()
	mx := mux.NewRouter()

	mx.Use(corsMiddleware)

	// Configure socket.io
	Api.InitializeRoutes(mx)

	if prod {
		logrus.Info("Starting Production Server...")
	} else {
		logrus.Info("Starting Development Server...")
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", env.PORT), mx))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request came from ", r.RemoteAddr)
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
