package main

import (
	"Assigment-Restfull_Customer/internal/controller"
	"Assigment-Restfull_Customer/internal/helper"
	repository "Assigment-Restfull_Customer/internal/repository/mysql"
	"Assigment-Restfull_Customer/internal/usecase"
	"Assigment-Restfull_Customer/middleware"
	"Assigment-Restfull_Customer/pkg"
	"Assigment-Restfull_Customer/router"

	"net/http"

	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := pkg.NewDB()
	validate := validator.New()
	customerRepository := repository.NewCustomerRepository()
	customerUsecase := usecase.NewCustomerUsecase(customerRepository, db, validate)

	customerController := controller.NewCustomerController(customerUsecase)

	router := router.NewRouter(customerController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	//server.ListenAndServe()

}
