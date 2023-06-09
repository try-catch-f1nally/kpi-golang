package services

import (
	"errors"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
)

type ReviewService struct {
	Db             *gorm.DB
	ProductService *ProductService
}

type CreateReviewBody struct {
	UserID    uint   `json:"userId"`
	ProductID uint   `json:"productId"`
	Rating    int    `json:"rating"`
	Text      string `json:"text"`
}

func (service *ReviewService) Create(reviewBody *CreateReviewBody) error {
	review := models.Review{
		UserID:    reviewBody.UserID,
		ProductID: reviewBody.ProductID,
		Rating:    reviewBody.Rating,
		Text:      reviewBody.Text,
	}
	err := service.Db.Create(&review).Error
	if err != nil {
		return err
	}
	return service.ProductService.RecalculateRating(reviewBody.ProductID)
}

func (service *ReviewService) Delete(reviewId uint, userId uint) error {
	var review models.Review
	err := service.Db.First(&review, reviewId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &utils.BadRequestError{Message: "wrong review ID"}
	}
	if err != nil {
		return err
	}
	if userId != review.UserID {
		return &utils.BadRequestError{Message: "deleting review of another user is not allowed"}
	}

	err = service.Db.Delete(&models.Review{}, reviewId).Error
	if err != nil {
		return err
	}

	return service.ProductService.RecalculateRating(review.ProductID)
}
