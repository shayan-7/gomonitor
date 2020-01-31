package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

var db gorm.DB

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

var key = []byte("sample_jwt_key")

type URLCredentials struct {
	Address   string `json:"address"`
	Threshold int    `json:"threshold"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenJson struct {
	Token string `json:"token"`
}

func setup() gorm.DB {
	db, err := gorm.Open(
		"postgres",
		"host=localhost user=postgres dbname=postgres password=postgres",
	)

	if err != nil {
		fmt.Printf("%+v\n", err)
		panic("Failed to connect to postgres DB")
	}
	var dbExists bool
	db.Raw(
		"SELECT EXISTS(SELECT * FROM pg_database WHERE datname = 'gomono');",
	).Row().Scan(&dbExists)

	if !dbExists {
		db.Exec("CREATE DATABASE gomono;")
	}
	db.Close()
	db, err = gorm.Open(
		"postgres",
		"host=localhost user=postgres dbname=gomono password=postgres",
	)
	db.AutoMigrate(&Member{})
	db.AutoMigrate(&URL{})
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
	http.HandleFunc("/apiv1/urls", handleURL)
	http.HandleFunc("/apiv1/alerts", getAlert)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
