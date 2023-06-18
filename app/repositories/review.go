package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/models"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (repo *ReviewRepository) ReviewCreate(review *models.Review) error {
	return repo.db.Create(review).Error
}

func (repo *ReviewRepository) ReviewGet(reviewId uint) (*models.Review, error) {
	var review models.Review
	err := repo.db.First(&review, reviewId).Error
	if err != nil {
		return nil, err
	}
	return &review, err
}

func (repo *ReviewRepository) ReviewGetByProductId(productId uint) ([]*models.Review, error) {
	var reviews []*models.Review
	err := repo.db.Where("product_id = ?", productId).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (repo *ReviewRepository) ReviewDelete(reviewId uint) error {
	return repo.db.Delete(&models.Review{}, reviewId).Error
}
