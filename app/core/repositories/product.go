package repositories

import "kpi-golang/app/core/models"

type ProductFilter struct {
	Types    []string
	Models   []string
	Memories []int
	Colors   []string
	Offset   int
	Limit    int
}

type ProductRepository interface {
	GetByIDs(productIDs []uint) ([]*models.Product, error)
	GetByFilter(productFilter *ProductFilter) ([]*models.Product, error)
	UpdateRating(productID uint, rating float64) error
}
