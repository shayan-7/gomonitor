package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Url struct {
	gorm.Model
	Address string
	Hash    string
}
