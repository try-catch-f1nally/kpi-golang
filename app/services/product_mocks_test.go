//go:build unit_test

package services

import "kpi-golang/app/models"

type ProductRepositoryMock struct {
	GetSuccess []*models.Product
	GetError   error
}

func (repo *ProductRepositoryMock) ProductGetByIds(productIds []uint) ([]*models.Product, error) {
	return repo.GetSuccess, repo.GetError
}

func (repo *ProductRepositoryMock) ProductGetByFilter(productFilter *ProductFilter) ([]*models.Product, error) {
	return repo.GetSuccess, repo.GetError
}

func (repo *ProductRepositoryMock) ProductUpdateRating(productId uint, rating float64) error {
	return repo.GetError
}
