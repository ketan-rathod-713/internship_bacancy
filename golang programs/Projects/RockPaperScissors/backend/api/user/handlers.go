package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rockpaperscissors/models"

	"github.com/sirupsen/logrus"
)

func (a *api) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (a *api) PostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post User Route Called")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := a.Service.CreateUser(&user)

	if err != nil {
		logrus.Error("error occured", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func (a *api) SignInUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	log.Print("signin request came")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := a.Service.SignInUser(&user); err != nil {
		logrus.Error("error occurred", err)
		http.Error(w, "failed to sign in user", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		log.Fatal(err)
	}
}
