package repositories

import "kpi-golang/app/core/models"

type ReviewRepository interface {
	Create(review *models.Review) error
	Get(reviewID uint) (*models.Review, error)
	GetByProductID(productID uint) ([]*models.Review, error)
	Delete(reviewID uint) error
}
