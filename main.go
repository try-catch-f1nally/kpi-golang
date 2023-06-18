package main

import (
	"github.com/gorilla/mux"
	"kpi-golang/app/controllers"
	"kpi-golang/app/db"
	"kpi-golang/app/repositories"
	"kpi-golang/app/services"
	"kpi-golang/app/utils"
	"log"
	"net/http"
)

func main() {
	database := db.Init()
	log.Println("Connected to DB")

	orderRepository := repositories.NewOrderRepository(database)
	productRepository := repositories.NewProductRepository(database)
	reviewRepository := repositories.NewReviewRepository(database)
	userRepository := repositories.NewUserRepository(database)

	tokenService := services.NewTokenService()
	authService := services.NewAuthService(userRepository, tokenService)
	orderService := services.NewOrderService(orderRepository, productRepository)
	productService := services.NewProductService(productRepository, reviewRepository)
	reviewService := services.NewReviewService(reviewRepository, productService)
	userService := services.NewUserService(userRepository)

	authController := controllers.NewAuthController(authService)
	orderController := controllers.NewOrderController(orderService, tokenService)
	productController := controllers.NewProductController(productService)
	reviewController := controllers.NewReviewController(reviewService, tokenService)
	userController := controllers.NewUserController(userService, tokenService)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	authController.RegisterRoutes(apiRouter)
	orderController.RegisterRoutes(apiRouter)
	productController.RegisterRoutes(apiRouter)
	reviewController.RegisterRoutes(apiRouter)
	userController.RegisterRoutes(apiRouter)

	port := utils.GetEnvVar("PORT", "3000")
	server := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:" + port,
	}
	log.Println("Application started listening on port " + port)
	log.Fatal(server.ListenAndServe())
}
