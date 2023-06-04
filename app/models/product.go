package models

import "gorm.io/gorm"

type GormModel struct {
	gorm.Model
}

type Product struct {
	GormModel
	Type   string
	Name   string
	Price  int
	Model  string
	Memory int
	Color  string
	Rating int
}
