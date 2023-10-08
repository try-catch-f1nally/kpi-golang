package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/core/models"
	"kpi-golang/app/core/repositories"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) GetByIDs(productIDs []uint) ([]*models.Product, error) {
	var products []*models.Product
	err := repo.db.Find(&products, productIDs).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) GetByFilter(productFilter *repositories.ProductFilter) ([]*models.Product, error) {
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

func (repo *ProductRepository) UpdateRating(productID uint, rating float64) error {
	return repo.db.Model(&models.Product{}).Where("id = ?", productID).Update("rating", rating).Error
}
