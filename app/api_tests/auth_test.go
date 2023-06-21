package api_tests

import (
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"kpi-golang/app/controllers"
	"kpi-golang/app/db"
	"kpi-golang/app/repositories"
	"kpi-golang/app/services"
	"log"
	"net/http"
	"testing"
)

func TestAll(t *testing.T) {
	database := db.Init()
	userRepository := repositories.NewUserRepository(database)
	tokenService := services.NewTokenService()
	authService := services.NewAuthService(userRepository, tokenService)
	authController := controllers.NewAuthController(authService)
	router := mux.NewRouter()
	authController.RegisterRoutes(router)

	err := database.Exec("TRUNCATE TABLE users").Error
	if err != nil {
		log.Fatal(err)
	}

	apitest.New("correct register").
		Handler(router).
		Post("/auth/register").
		Body(`{"first_name":"name","last_name":"surname","email":"name@mail.com","password":"12345pwd"}`).
		Expect(t).
		Assert(jsonpath.Present("$.accessToken")).
		Assert(jsonpath.Present("$.refreshToken")).
		Assert(jsonpath.Present("$.userId")).
		Status(http.StatusCreated).
		End()

	apitest.New("test register with existing email").
		Handler(router).
		Post("/auth/register").
		Body(`{"first_name":"name2","last_name":"surname2","email":"name@mail.com","password":"54321pwd"}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	apitest.New("correct login").
		Handler(router).
		Post("/auth/login").
		Body(`{"email":"name@mail.com","password":"12345pwd"}`).
		Expect(t).
		Assert(jsonpath.Present("$.accessToken")).
		Assert(jsonpath.Present("$.refreshToken")).
		Assert(jsonpath.Present("$.userId")).
		Status(http.StatusOK).
		End()

	apitest.New("login with incorrect email").
		Handler(router).
		Post("/auth/login").
		Body(`{"email":"wrong@mail.com","password":"12345pwd"}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	apitest.New("login with incorrect password").
		Handler(router).
		Post("/auth/login").
		Body(`{"email":"name@mail.com","password":"wrong"}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
