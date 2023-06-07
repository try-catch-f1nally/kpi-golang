package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kpi-golang/app/services"
	"kpi-golang/app/utils"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService  *services.UserService
	TokenService *services.TokenService
}

func (controller *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/{id}/first-name", controller.changeFirstName).Methods("PATCH")
	router.HandleFunc("/users/{id}/last-name", controller.changeLastName).Methods("PATCH")
	router.HandleFunc("/users/{id}/email", controller.changeEmail).Methods("PATCH")
	router.HandleFunc("/users/{id}/password", controller.changePassword).Methods("PATCH")
}

func (controller *UserController) changeFirstName(w http.ResponseWriter, r *http.Request) {
	userIdInToken, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	userId := parseUserIdFromParams(w, r)
	if userId != userIdInToken {
		http.Error(w, "changing the first name of another user is not allowed", http.StatusForbidden)
		return
	}

	var body struct {
		FirstName string `json:"firstName"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	err = controller.UserService.ChangeFirstName(userId, body.FirstName)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (controller *UserController) changeLastName(w http.ResponseWriter, r *http.Request) {
	userIdInToken, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	userId := parseUserIdFromParams(w, r)
	if userId != userIdInToken {
		http.Error(w, "changing the last name of another user is not allowed", http.StatusForbidden)
		return
	}

	var body struct {
		LastName string `json:"lastName"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	err = controller.UserService.ChangeLastName(userId, body.LastName)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (controller *UserController) changeEmail(w http.ResponseWriter, r *http.Request) {
	userIdInToken, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	userId := parseUserIdFromParams(w, r)
	if userId != userIdInToken {
		http.Error(w, "changing the email of another user is not allowed", http.StatusForbidden)
		return
	}

	var body struct {
		Email string `json:"email"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	err = controller.UserService.ChangeEmail(userId, body.Email)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (controller *UserController) changePassword(w http.ResponseWriter, r *http.Request) {
	userIdInToken, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	userId := parseUserIdFromParams(w, r)
	if userId != userIdInToken {
		http.Error(w, "changing the password of another user is not allowed", http.StatusForbidden)
		return
	}

	var body struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
	}

	err = controller.UserService.ChangePassword(userId, body.OldPassword, body.NewPassword)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseUserIdFromParams(w http.ResponseWriter, r *http.Request) uint {
	userId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, "wrong user ID", http.StatusBadRequest)
	}
	return uint(userId)
}
