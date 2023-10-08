package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/core/models"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (repo *ReviewRepository) Create(review *models.Review) error {
	return repo.db.Create(review).Error
}

func (repo *ReviewRepository) Get(reviewID uint) (*models.Review, error) {
	var review models.Review
	err := repo.db.First(&review, reviewID).Error
	if err != nil {
		return nil, err
	}
	return &review, err
}

func (repo *ReviewRepository) GetByProductID(productID uint) ([]*models.Review, error) {
	var reviews []*models.Review
	err := repo.db.Where("product_id = ?", productID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (repo *ReviewRepository) Delete(reviewID uint) error {
	return repo.db.Delete(&models.Review{}, reviewID).Error
}
