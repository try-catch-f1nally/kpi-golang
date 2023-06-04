package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   string
	Number   int `gorm:"unique;autoIncrement"`
	Payment  string
	Delivery string
	Products []Product `gorm:"many2many:order_products"`
}
