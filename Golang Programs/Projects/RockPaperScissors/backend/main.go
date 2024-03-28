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
	mux := mux.NewRouter()

	mux.Use(corsMiddleware)

	// Configure socket.io
	Api.InitializeRoutes(mux)

	if prod {
		logrus.Info("Starting Production Server...")
	} else {
		logrus.Info("Starting Development Server...")
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", env.PORT), mux))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		// w.Header().Set("Access-Control-Allow-Origin", "http://192.168.7.33:5500")
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
