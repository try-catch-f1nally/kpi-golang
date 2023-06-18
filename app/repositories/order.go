package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (repo *OrderRepository) OrderCreate(order *models.Order) error {
	return repo.db.Create(order).Error
}

func (repo *OrderRepository) OrderGetByUserId(userId uint) ([]*models.Order, error) {
	var orders []*models.Order
	err := repo.db.Preload("Products").Where("user_id = ?", userId).Order("created_at desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
