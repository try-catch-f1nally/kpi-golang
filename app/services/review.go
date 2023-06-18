package services

import (
	"errors"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
)

type ReviewRepository interface {
	ReviewCreate(review *models.Review) error
	ReviewGet(reviewId uint) (*models.Review, error)
	ReviewGetByProductId(productId uint) ([]*models.Review, error)
	ReviewDelete(reviewId uint) error
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
	err := service.reviewRepository.ReviewCreate(&models.Review{
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

func (service *ReviewService) Delete(reviewId uint, userId uint) error {
	review, err := service.reviewRepository.ReviewGet(reviewId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &utils.BadRequestError{Message: "wrong review ID"}
	}
	if err != nil {
		return err
	}
	if userId != review.UserID {
		return &utils.BadRequestError{Message: "deleting review of another user is not allowed"}
	}

	err = service.reviewRepository.ReviewDelete(reviewId)
	if err != nil {
		return err
	}

	return service.productService.RecalculateRating(review.ProductID)
}
