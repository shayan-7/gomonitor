package main

import (
	"fmt"
	//"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"log"
	"net/http"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

var key = []byte("sample_jwt_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWT struct {
	Username string `json:"token"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	creds := Credentials{}
	fmt.Printf("\033[36;1m%s, %s", creds.Username, creds.Password)

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password, found := users[creds.Username]
	if !found || password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("\033[92mServing on localhost:8090")
	http.HandleFunc("/login", Login)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
