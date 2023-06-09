package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint       `json:"userId"`
	Number   uint       `json:"number" gorm:"unique;autoIncrement"`
	Payment  string     `json:"payment"`
	Delivery string     `json:"delivery"`
	Products []*Product `json:"products" gorm:"many2many:order_products"`
}
