package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kpi-golang/app/services"
	"kpi-golang/app/utils"
	"net/http"
)

type UserController struct {
	UserService  *services.UserService
	TokenService *services.TokenService
}

func (controller *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/current/first-name", controller.changeFirstName).Methods("PATCH")
	router.HandleFunc("/users/current/last-name", controller.changeLastName).Methods("PATCH")
	router.HandleFunc("/users/current/email", controller.changeEmail).Methods("PATCH")
	router.HandleFunc("/users/current/password", controller.changePassword).Methods("PATCH")
}

func (controller *UserController) changeFirstName(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
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
	userId, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
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
	userId, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
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
	userId, err := utils.AuthenticateRequest(r, controller.TokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	var body struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	err = controller.UserService.ChangePassword(userId, body.OldPassword, body.NewPassword)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
