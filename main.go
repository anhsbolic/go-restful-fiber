package main

import (
	"github.com/go-playground/validator/v10"
	"go-restful-fiber/app"
	"go-restful-fiber/controller"
	"go-restful-fiber/helper"
	"go-restful-fiber/middleware"
	"go-restful-fiber/repository"
	"go-restful-fiber/service"
	"net/http"
)

func main() {
	// Setup DB
	db := app.NewDB()

	// Setup Validator
	validate := validator.New()

	// Setup Category API
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	// Setup Server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
