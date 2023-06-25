//go:build unit_test

package services

import "kpi-golang/app/models"

type OrderRepositoryMock struct {
	GetSuccess  []*models.Order
	GetError    error
	CreateError error
}

func (repo *OrderRepositoryMock) OrderGetByUserId(userId uint) ([]*models.Order, error) {
	return repo.GetSuccess, repo.GetError
}

func (repo *OrderRepositoryMock) OrderCreate(order *models.Order) error {
	return repo.CreateError
}
