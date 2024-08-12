package main

import (
	"app/routinity/app"
	"app/routinity/controller"
	"app/routinity/helper"
	"app/routinity/middleware"
	"app/routinity/repository"
	"app/routinity/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	routinityRepository := repository.NewRoutintityRepo()
	routinityService := service.NewRoutinityService(routinityRepository, db, validate)
	routinityController := controller.NewRoutinityController(routinityService)
	router := app.NewRouter(routinityController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}