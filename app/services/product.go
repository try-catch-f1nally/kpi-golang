package services

import (
	"kpi-golang/app/models"
)

type ProductRepository interface {
	ProductGetByIds(productIds []uint) ([]*models.Product, error)
	ProductGetByFilter(productFilter *ProductFilter) ([]*models.Product, error)
	ProductUpdateRating(productId uint, rating float64) error
}

type ProductService struct {
	productRepository ProductRepository
	reviewRepository  ReviewRepository
}

func NewProductService(productRepository ProductRepository, reviewRepository ReviewRepository) *ProductService {
	return &ProductService{productRepository, reviewRepository}
}

type ProductFilter struct {
	Types    []string
	Models   []string
	Memories []int
	Colors   []string
	Offset   int
	Limit    int
}

func (service *ProductService) GetProducts(productFilter *ProductFilter) ([]*models.Product, error) {
	products, err := service.productRepository.ProductGetByFilter(productFilter)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service *ProductService) RecalculateRating(productId uint) error {
	reviews, err := service.reviewRepository.ReviewGetByProductId(productId)
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
	return service.productRepository.ProductUpdateRating(productId, rating)
}
