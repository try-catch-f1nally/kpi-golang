package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kpi-golang/app/services"
	"kpi-golang/app/utils"
	"net/http"
)

type ReviewController struct {
	reviewService *services.ReviewService
	tokenService  *services.TokenService
}

func NewReviewController(reviewService *services.ReviewService, tokenService *services.TokenService) *ReviewController {
	return &ReviewController{reviewService, tokenService}
}

func (controller *ReviewController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/reviews", controller.postReview).Methods("POST")
	router.HandleFunc("/reviews/{id}", controller.deleteReview).Methods("DELETE")
}

func (controller *ReviewController) postReview(w http.ResponseWriter, r *http.Request) {
	userIdInToken, err := utils.AuthenticateRequest(r, controller.tokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	var body services.CreateReviewBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	if userIdInToken != body.UserID {
		http.Error(w, "posting review for another user is not allowed", http.StatusBadRequest)
	}

	err = controller.reviewService.Create(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (controller *ReviewController) deleteReview(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.AuthenticateRequest(r, controller.tokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	reviewId, err := utils.ParseEntityIdFromParams(r)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	err = controller.reviewService.Delete(reviewId, userId)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
