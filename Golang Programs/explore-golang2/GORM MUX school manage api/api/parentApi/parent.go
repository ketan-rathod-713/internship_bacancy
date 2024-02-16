package parentapi

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"schoolApi/models"
)

var apiLogs = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var errorLogs = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func (a *api) getAllParents(w http.ResponseWriter, r *http.Request) {

}

func (a *api) createParent(w http.ResponseWriter, r *http.Request) {
	parent := &models.Parent{}
	err := json.NewDecoder(r.Body).Decode(parent)

	if err != nil {
		errorLogs.Println("Error Fetching Json Data", err)
		// send json eror here
	}

	parent, err = a.Service.CreateParent(parent)
	if err != nil {
		http.Error(w, "error occured", http.DefaultMaxHeaderBytes)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parent)
}
