package repositories

import "kpi-golang/app/core/models"

type OrderRepository interface {
	Create(order *models.Order) error
	GetByUserId(userID uint) ([]*models.Order, error)
}
