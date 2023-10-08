package services

import (
	"kpi-golang/app/core/models"
	"kpi-golang/app/core/repositories"
)

type ProductService struct {
	productRepository repositories.ProductRepository
	reviewRepository  ReviewRepository
}

func NewProductService(productRepository repositories.ProductRepository, reviewRepository ReviewRepository) *ProductService {
	return &ProductService{productRepository, reviewRepository}
}

func (service *ProductService) GetProducts(productFilter *repositories.ProductFilter) ([]*models.Product, error) {
	products, err := service.productRepository.GetByFilter(productFilter)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service *ProductService) RecalculateRating(productID uint) error {
	reviews, err := service.reviewRepository.GetByProductID(productID)
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
	return service.productRepository.UpdateRating(productID, rating)
}
