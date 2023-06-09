package models

import "gorm.io/gorm"

type GormModel struct {
	gorm.Model
}

type Product struct {
	GormModel
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Price   int      `json:"price"`
	Model   string   `json:"model"`
	Memory  int      `json:"memory"`
	Color   string   `json:"color"`
	Rating  float64  `json:"rating"`
	Reviews []Review `json:"reviews"`
}
