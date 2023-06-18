package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kpi-golang/app/services"
	"kpi-golang/app/utils"
	"net/http"
)

type OrderController struct {
	orderService *services.OrderService
	tokenService *services.TokenService
}

func NewOrderController(orderService *services.OrderService, tokenService *services.TokenService) *OrderController {
	return &OrderController{orderService, tokenService}
}

func (controller *OrderController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orders", controller.getOrders).Methods("GET")
	router.HandleFunc("/orders", controller.createOrder).Methods("POST")
}

func (controller *OrderController) getOrders(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.AuthenticateRequest(r, controller.tokenService.ValidateAccessToken)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	orders, err := controller.orderService.GetOrders(userId)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	utils.RespondWithJson(w, orders)
}

func (controller *OrderController) createOrder(w http.ResponseWriter, r *http.Request) {
	var body services.CreateOrderBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	err = controller.orderService.CreateOrder(&body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
