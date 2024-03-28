package user

import (
	"encoding/json"
	"net/http"
	"rockpaperscissors/models"

	"github.com/sirupsen/logrus"
)

func (a *api) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (a *api) PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := a.Service.CreateUser(&user)

	if err != nil {
		logrus.Error("error occured", err)
	}

	json.NewEncoder(w).Encode(&user)
}

func (a *api) SignInUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := a.Service.SignInUser(&user)

	if err != nil {
		logrus.Error("error occured", err)
	}

	json.NewEncoder(w).Encode(&user)
}
