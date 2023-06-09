package services

import (
	"errors"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
)

type OrderService struct {
	Db *gorm.DB
}

type CreateOrderBody struct {
	UserID     uint   `json:"userId"`
	Payment    string `json:"payment"`
	Delivery   string `json:"delivery"`
	ProductIds []uint `json:"productIds"`
}

func (service *OrderService) GetOrders(userId uint) ([]models.Order, error) {
	var orders []models.Order
	err := service.Db.Preload("Products").Where("user_id = ?", userId).Order("created_at desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (service *OrderService) CreateOrder(createOrderBody *CreateOrderBody) error {
	var products []*models.Product
	err := service.Db.Find(&products, createOrderBody.ProductIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &utils.BadRequestError{Message: "wrong product ID"}
	}
	if err != nil {
		return err
	}

	order := models.Order{
		UserID:   createOrderBody.UserID,
		Delivery: createOrderBody.Delivery,
		Products: products,
	}
	return service.Db.Model(&models.Order{}).Create(&order).Error
}
