package sportapi

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"schoolApi/models"
)

var apiLogs = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var errorLogs = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func (a *api) getAllSports(w http.ResponseWriter, r *http.Request) {

}

func (a *api) createSport(w http.ResponseWriter, r *http.Request) {
	sport := &models.Sport{}
	err := json.NewDecoder(r.Body).Decode(sport)

	if err != nil {
		errorLogs.Println("Error Fetching Json Data", err)
		// send json eror here
	}

	sport, err = a.Service.CreateSport(sport)
	if err != nil {
		http.Error(w, "error occured", http.DefaultMaxHeaderBytes)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sport)
}
