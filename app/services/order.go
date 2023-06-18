package services

import (
	"errors"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
)

type OrderRepository interface {
	OrderCreate(order *models.Order) error
	OrderGetByUserId(userId uint) ([]*models.Order, error)
}

type OrderService struct {
	orderRepository   OrderRepository
	productRepository ProductRepository
}

func NewOrderService(orderRepository OrderRepository, productRepository ProductRepository) *OrderService {
	return &OrderService{orderRepository, productRepository}
}

type CreateOrderBody struct {
	UserID     uint   `json:"userId"`
	Payment    string `json:"payment"`
	Delivery   string `json:"delivery"`
	ProductIds []uint `json:"productIds"`
}

func (service *OrderService) GetOrders(userId uint) ([]*models.Order, error) {
	orders, err := service.orderRepository.OrderGetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (service *OrderService) CreateOrder(createOrderBody *CreateOrderBody) error {
	products, err := service.productRepository.ProductGetByIds(createOrderBody.ProductIds)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &utils.BadRequestError{Message: "wrong product ID"}
	}
	if err != nil {
		return err
	}

	return service.orderRepository.OrderCreate(&models.Order{
		UserID:   createOrderBody.UserID,
		Delivery: createOrderBody.Delivery,
		Products: products,
	})
}
