package controllers

import (
	"github.com/gorilla/mux"
	"kpi-golang/app/services"
	"kpi-golang/app/utils"
	"net/http"
	"strconv"
	"strings"
)

type ProductController struct {
	ProductService *services.ProductService
}

func (controller *ProductController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", controller.getProducts).Methods("GET")
}

func (controller *ProductController) getProducts(w http.ResponseWriter, r *http.Request) {
	productFilter := &services.ProductFilter{}

	types := r.URL.Query().Get("type")
	if types != "" {
		productFilter.Types = strings.Split(types, ",")
	}

	models := r.URL.Query().Get("model")
	if models != "" {
		productFilter.Models = strings.Split(models, ",")
	}

	memories := r.URL.Query().Get("memory")
	if memories != "" {
		memoriesList := strings.Split(memories, ",")
		for _, mem := range memoriesList {
			memInt, err := strconv.Atoi(mem)
			if err != nil {
				http.Error(w, "memory values should be integers", http.StatusBadRequest)
				return
			}
			productFilter.Memories = append(productFilter.Memories, memInt)
		}
	}

	colors := r.URL.Query().Get("color")
	if colors != "" {
		productFilter.Colors = strings.Split(colors, ",")
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, "offset value should be an integer", http.StatusBadRequest)
			return
		}
		productFilter.Offset = offset
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "limit value should be an integer", http.StatusBadRequest)
			return
		}
		productFilter.Limit = limit
	}

	products, err := controller.ProductService.GetProducts(productFilter)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJson(w, products)
}
