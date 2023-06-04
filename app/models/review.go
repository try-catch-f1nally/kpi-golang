package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    string
	ProductID string
	Rating    int
	Text      string
}
