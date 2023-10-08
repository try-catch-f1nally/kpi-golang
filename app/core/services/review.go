package services

import (
	"errors"
	"gorm.io/gorm"
	"kpi-golang/app/core"
	"kpi-golang/app/core/models"
)

type ReviewRepository interface {
	Create(review *models.Review) error
	Get(reviewID uint) (*models.Review, error)
	GetByProductID(productID uint) ([]*models.Review, error)
	Delete(reviewID uint) error
}

type ReviewService struct {
	reviewRepository ReviewRepository
	productService   *ProductService
}

func NewReviewService(reviewRepository ReviewRepository, productService *ProductService) *ReviewService {
	return &ReviewService{reviewRepository, productService}
}

type CreateReviewBody struct {
	UserID    uint   `json:"userId"`
	ProductID uint   `json:"productId"`
	Rating    int    `json:"rating"`
	Text      string `json:"text"`
}

func (service *ReviewService) Create(reviewBody *CreateReviewBody) error {
	err := service.reviewRepository.Create(&models.Review{
		UserID:    reviewBody.UserID,
		ProductID: reviewBody.ProductID,
		Rating:    reviewBody.Rating,
		Text:      reviewBody.Text,
	})
	if err != nil {
		return err
	}
	return service.productService.RecalculateRating(reviewBody.ProductID)
}

func (service *ReviewService) Delete(reviewID uint, userID uint) error {
	review, err := service.reviewRepository.Get(reviewID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &core.BadRequestError{Message: "wrong review ID"}
	}
	if err != nil {
		return err
	}
	if userID != review.UserID {
		return &core.BadRequestError{Message: "deleting review of another user is not allowed"}
	}

	err = service.reviewRepository.Delete(reviewID)
	if err != nil {
		return err
	}

	return service.productService.RecalculateRating(review.ProductID)
}
