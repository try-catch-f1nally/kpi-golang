package services

import (
	"gorm.io/gorm"
	"kpi-golang/app/models"
)

type ProductService struct {
	Db *gorm.DB
}

type ProductFilter struct {
	Types    []string
	Models   []string
	Memories []int
	Colors   []string
	Offset   int
	Limit    int
}

func (service *ProductService) GetProducts(productFilter *ProductFilter) ([]models.Product, error) {
	var products []models.Product
	db := service.Db.Preload("Reviews")

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

func (service *ProductService) RecalculateRating(productId uint) error {
	var reviews []models.Review
	err := service.Db.Where("product_id = ?", productId).Find(&reviews).Error
	if err != nil {
		return err
	}
	rating := 0.0
	if len(reviews) > 0 {
		sum := 0.0
		for _, el := range reviews {
			sum += float64(el.Rating)
		}
		rating = sum / float64(len(reviews))
	}
	err = service.Db.Model(&models.Product{}).Where("id = ?", productId).Update("rating", rating).Error
	return err
}
