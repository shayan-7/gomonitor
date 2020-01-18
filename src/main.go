package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"time"
)

var db gorm.DB

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

var key = []byte("sample_jwt_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenJson struct {
	Token string `json:"token"`
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
	response, _ := json.MarshalIndent(
		member,
		"",
		"  ",
	)
	w.Header().Set("Content-Type", "application/json")
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

func setup() gorm.DB {
	db, err := gorm.Open(
		"sqlite3",
		"gomonitor.db",
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic("Failed to connect to postgres DB")
	}
	db.AutoMigrate(&Member{})

	return *db
}

func teardown(db gorm.DB) {
	db.Close()
}

func main() {
	db = setup()
	defer teardown(db)

	fmt.Println("\033[92mServing on localhost:8090")
	http.HandleFunc("/apiv1/tokens", createToken)
	http.HandleFunc("/apiv1/members", registerMember)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
