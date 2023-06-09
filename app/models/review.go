package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    uint   `json:"userId"`
	ProductID uint   `json:"productId"`
	Rating    int    `json:"rating"`
	Text      string `json:"text"`
}
