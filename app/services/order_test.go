//go:build unit_test

package services

import (
	"errors"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"reflect"
	"testing"
)

func TestGetOrders(t *testing.T) {
	orderGetError := errors.New("failed to get orders")
	var userId uint = 21
	order := &models.Order{
		UserID:   userId,
		Payment:  "online; google pay",
		Delivery: "212 Pearl St, New York, NY 10038, United States",
	}

	testCases := []struct {
		description     string
		orderGetSuccess []*models.Order
		orderGetError   error
		expectedSuccess []*models.Order
		expectedError   error
	}{
		{
			description:   "get error from order repository",
			orderGetError: orderGetError,
			expectedError: orderGetError,
		},
		{
			description:   "not found error from order repository",
			orderGetError: gorm.ErrRecordNotFound,
			expectedError: gorm.ErrRecordNotFound,
		},
		{
			description:     "success order, empty order product list",
			orderGetSuccess: []*models.Order{order},
			expectedSuccess: []*models.Order{order},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			service := NewOrderService(
				&OrderRepositoryMock{GetSuccess: tc.orderGetSuccess, GetError: tc.orderGetError},
				&ProductRepositoryMock{})

			result, err := service.GetOrders(userId)
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
			if !reflect.DeepEqual(tc.expectedSuccess, result) {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedSuccess, result)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	orderCreateError := errors.New("failed to create orders")
	testCases := []struct {
		description       string
		orderCreateError  error
		productGetSuccess []*models.Product
		orderGetError     error
		productGetError   error
		expectedSuccess   []*models.Order
		expectedError     error
	}{
		{
			description:      "get error from order repository",
			orderCreateError: orderCreateError,
			expectedError:    orderCreateError,
		},
		{
			description: "success order create",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			service := NewOrderService(
				&OrderRepositoryMock{CreateError: tc.orderCreateError, GetError: tc.orderGetError},
				&ProductRepositoryMock{GetSuccess: tc.productGetSuccess, GetError: tc.productGetError})

			err := service.CreateOrder(&CreateOrderBody{
				UserID:   21,
				Payment:  "online; google pay",
				Delivery: "212 Pearl St, New York, NY 10038, United States",
			})
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}
