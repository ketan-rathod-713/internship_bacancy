package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

// map of username and password
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// claims // payload for our jwt
type Claims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // extra information about jwt.
	// ffor now we will use ExpiresAt variable for now
}

// Username and password is coming
// if any error then BAD REQUEST
// then created claims object
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println("Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// lets match our data
	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		log.Println("Unauthorised")
		w.WriteHeader(http.StatusUnauthorized) // This request is not authorised
		return
	}

	// if req is authorised then create claims object and send it back with token
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{expirationTime},
		},
	}

	// create token from it

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// from this token get token string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//get value of token
	tokenStr := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// claims me claim aa jaega

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid { // token is not valid so return status invalid
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello %s", claims.Username)))
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//get value of token
	tokenStr := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// claims me claim aa jaega

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid { // token is not valid so return status invalid
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TILL here we have done same as home handler

	// jab 30 seconda baki hoge tab ham refresh token generate karege
	// TODO: Sub, Unix
	// if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
	// w.WriteHeader(http.StatusBadRequest)
	// return
	// }

	// copy from login handler
	expirationTime := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = &jwt.NumericDate{expirationTime}

	// create token from it

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// from this token get token string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token", // TODO: Other wise same token name can be used
		Value:   tokenString,
		Expires: expirationTime,
	})
}
