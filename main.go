package main

import (
	"github.com/gorilla/mux"
	"kpi-golang/app/controllers"
	"kpi-golang/app/db"
	"kpi-golang/app/services"
	"kpi-golang/app/utils"
	"log"
	"net/http"
)

func main() {
	database := db.Init()
	log.Println("Connected to DB")

	tokenService := &services.TokenService{}
	authService := &services.AuthService{Db: database, TokenService: tokenService}
	//orderService := &services.OrderService{Db: database}
	productService := &services.ProductService{Db: database}
	//reviewService := &services.ReviewService{Db: database}
	userService := &services.UserService{Db: database}

	authController := &controllers.AuthController{AuthService: authService}
	//orderController := controllers.OrderController{OrderService: orderService}
	productController := &controllers.ProductController{ProductService: productService}
	//reviewController := controllers.ReviewController{ReviewService: reviewService}
	userController := &controllers.UserController{UserService: userService, TokenService: tokenService}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	authController.RegisterRoutes(apiRouter)
	//orderController.RegisterRoutes(apiRouter)
	productController.RegisterRoutes(apiRouter)
	//reviewController.RegisterRoutes(apiRouter)
	userController.RegisterRoutes(apiRouter)

	port := utils.GetEnvVar("PORT", "3000")
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
	}
	log.Println("Application started listening on port " + port)
	log.Fatal(server.ListenAndServe())
}
