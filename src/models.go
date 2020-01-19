package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Member struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	//jwt.StandardClaims
}

type URL struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Address   string    `json:"address"`
	Failure   int       `json:"failure"`
	Success   int       `json:"success"`
	Threshold int       `json:"threshold"`
}
