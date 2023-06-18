package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/services"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) ProductGetByIds(productIds []uint) ([]*models.Product, error) {
	var products []*models.Product
	err := repo.db.Find(&products, productIds).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) ProductGetByFilter(productFilter *services.ProductFilter) ([]*models.Product, error) {
	var products []*models.Product
	db := repo.db.Preload("Reviews")

	if productFilter != nil {
		if len(productFilter.Types) > 0 {
			db = db.Where("type IN ?", productFilter.Types)
		}
		if len(productFilter.Models) > 0 {
			db = db.Where("model IN ?", productFilter.Models)
		}
		if len(productFilter.Memories) > 0 {
			db = db.Where("memory IN ?", productFilter.Memories)
		}
		if len(productFilter.Colors) > 0 {
			db = db.Where("color IN ?", productFilter.Colors)
		}
		if productFilter.Offset > 0 {
			db = db.Offset(productFilter.Offset)
		}
		if productFilter.Limit > 0 {
			db = db.Limit(productFilter.Limit)
		}
	}

	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) ProductUpdateRating(productId uint, rating float64) error {
	return repo.db.Model(&models.Product{}).Where("id = ?", productId).Update("rating", rating).Error
}
