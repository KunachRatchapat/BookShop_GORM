package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string 
	Lastname string
	Email    string `gorm:"unique"`
	Password string
}
