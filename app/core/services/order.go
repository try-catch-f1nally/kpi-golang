package services

import (
	"errors"
	"gorm.io/gorm"
	"kpi-golang/app/core"
	"kpi-golang/app/core/models"
	"kpi-golang/app/core/repositories"
)

type OrderService struct {
	orderRepository   repositories.OrderRepository
	productRepository repositories.ProductRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, productRepository repositories.ProductRepository) *OrderService {
	return &OrderService{orderRepository, productRepository}
}

type CreateOrderBody struct {
	UserID     uint   `json:"userId"`
	Payment    string `json:"payment"`
	Delivery   string `json:"delivery"`
	ProductIds []uint `json:"productIDs"`
}

func (service *OrderService) GetOrders(userID uint) ([]*models.Order, error) {
	orders, err := service.orderRepository.GetByUserId(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (service *OrderService) CreateOrder(createOrderBody *CreateOrderBody) error {
	products, err := service.productRepository.GetByIDs(createOrderBody.ProductIds)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &core.BadRequestError{Message: "wrong product ID"}
	}
	if err != nil {
		return err
	}

	return service.orderRepository.Create(&models.Order{
		UserID:   createOrderBody.UserID,
		Delivery: createOrderBody.Delivery,
		Products: products,
	})
}
