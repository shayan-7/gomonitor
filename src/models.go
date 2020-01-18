package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Member struct {
	gorm.Model
	Username string
	Password string `json:"-"`
}
