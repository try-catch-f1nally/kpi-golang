//go:build integration_test

package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/db"
	"kpi-golang/app/models"
	"log"
	"testing"
)

func cleanOrderTable(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE orders").Error
	if err != nil {
		log.Fatal(err)
	}
}

func TestOrderCreateGetByUserId(t *testing.T) {
	database := db.Init()
	cleanOrderTable(database)
	repo := NewOrderRepository(database)

	newOrder := &models.Order{
		UserID:   21,
		Payment:  "online; google pay",
		Delivery: "212 Pearl St, New York, NY 10038, United States",
	}

	t.Run("create and get by user ID new order", func(t *testing.T) {
		err := repo.OrderCreate(newOrder)
		if err != nil {
			t.Errorf("failed to create order with error: %v", err)
			return
		}

		orders, err := repo.OrderGetByUserId(newOrder.UserID)
		if err != nil {
			t.Errorf("failed to get order with error: %v", err)
			return
		}
		if len(orders) != 1 {
			t.Errorf("wrong orders amount returned; actual: %v, expected: 1", len(orders))
		}

		order := orders[0]
		if order.ID != newOrder.ID || order.Number < 1 || order.Payment != newOrder.Payment ||
			order.Delivery != newOrder.Delivery {
			t.Errorf("order data is corrupted; actual: %v, expected: %v", order, newOrder)
		}
	})
}
