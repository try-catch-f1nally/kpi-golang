package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"kpi-golang/app/core/services"
	"kpi-golang/app/utils"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (controller *AuthController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/register", controller.register).Methods("POST")
	router.HandleFunc("/auth/login", controller.login).Methods("POST")
	router.HandleFunc("/auth/logout", controller.logout).Methods("POST")
	router.HandleFunc("/auth/refresh", controller.refresh).Methods("POST")
}

func (controller *AuthController) register(w http.ResponseWriter, r *http.Request) {
	var body services.RegisterBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	userData, err := controller.authService.Register(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	setTokenInCookie(w, userData.RefreshToken)
	utils.RespondWithJson(w, userData)
}

func (controller *AuthController) login(w http.ResponseWriter, r *http.Request) {
	var body services.LoginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	userData, err := controller.authService.Login(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	setTokenInCookie(w, userData.RefreshToken)
	utils.RespondWithJson(w, userData)
}

func (controller *AuthController) logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshToken")
	if errors.Is(err, http.ErrNoCookie) {
		http.Error(w, "refresh token is missing", http.StatusUnauthorized)
		return
	}
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	unsetTokenInCookie(w)

	err = controller.authService.Logout(cookie.Value)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
}

func (controller *AuthController) refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshToken")
	if errors.Is(err, http.ErrNoCookie) {
		http.Error(w, "refresh token is missing", http.StatusUnauthorized)
		return
	}
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	unsetTokenInCookie(w)

	userData, err := controller.authService.Refresh(cookie.Value)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	setTokenInCookie(w, userData.RefreshToken)
	utils.RespondWithJson(w, userData)
}

func setTokenInCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    token,
		MaxAge:   30 * 24 * 60 * 60,
		HttpOnly: true,
		Path:     "/api/auth",
	})
}

func unsetTokenInCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "refreshToken",
		Value:  "",
		MaxAge: -1,
	})
}
