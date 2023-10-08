package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  []byte
	FirstName string
	LastName  string
	Token     string
}
