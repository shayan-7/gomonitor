package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func handleURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createURL(w, r)
	case "GET":
		getURL(w, r)
	}
}

func getURL(w http.ResponseWriter, r *http.Request) {
	isValid := isValidRequest(r)
	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	urls := []URL{}
	db.Find(&urls)
	response, _ := json.MarshalIndent(
		&urls,
		"",
		"  ",
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func isValidRequest(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return false
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)

	if err != nil || err == jwt.ErrSignatureInvalid || !token.Valid {
		return false
	}

	return true
}

func createURL(w http.ResponseWriter, r *http.Request) {
	isValid := isValidRequest(r)
	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	urlcreds := URLCredentials{}
	err := json.NewDecoder(r.Body).Decode(&urlcreds)
	if err != nil || urlcreds.Threshold == 0 {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url := URL{Address: urlcreds.Address, Threshold: urlcreds.Threshold}
	db.Create(&url)
	response, _ := json.MarshalIndent(
		&url,
		"",
		"  ",
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func registerMember(w http.ResponseWriter, r *http.Request) {
	creds := Credentials{}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	member := Member{}
	var membersCount int
	db.Where(
		"username = ?",
		creds.Username,
	).Find(&member).Count(&membersCount)
	if membersCount > 0 {
		w.WriteHeader(http.StatusConflict)
		return
	}

	member.Username = creds.Username
	member.Password = creds.Password
	db.Create(&member)

	claims := &Claims{
		ID:             member.ID,
		Username:       member.Username,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.Token{
		Header: map[string]interface{}{
			"alg": jwt.SigningMethodHS256.Alg(),
			"exp": time.Now().Add((time.Hour * 24) * 7).Unix(),
		},
		Claims: claims,
		Method: jwt.SigningMethodHS256,
	}
	tokenstring, err := token.SignedString(key)
	response, _ := json.MarshalIndent(
		member,
		"",
		"  ",
	)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-JWT-Token", tokenstring)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func createToken(w http.ResponseWriter, r *http.Request) {
	creds := Credentials{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(creds)

	member := Member{}
	var membersCount int
	db.Where(
		"username = ? AND password = ?",
		creds.Username,
		creds.Password,
	).Find(&member).Count(&membersCount)
	if membersCount == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expiredAt := time.Now().Add((time.Hour * 24) * 7)
	claims := &Claims{
		Username:       creds.Username,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.Token{
		Header: map[string]interface{}{
			"alg": jwt.SigningMethodHS256.Alg(),
			"exp": expiredAt.Unix(),
		},
		Claims: claims,
		Method: jwt.SigningMethodHS256,
	}
	tokenstring, err := token.SignedString(key)
	response, _ := json.MarshalIndent(
		&TokenJson{Token: tokenstring},
		"",
		"  ",
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
