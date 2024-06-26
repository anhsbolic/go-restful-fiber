package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-restful-fiber/app"
	"go-restful-fiber/config"
	"go-restful-fiber/controller"
	"go-restful-fiber/middleware"
	"go-restful-fiber/pkg"
	"go-restful-fiber/repository"
	"go-restful-fiber/service"
	"net/http"
)

func main() {
	// Get Config
	env := config.GetEnvConfig()

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
	addr := fmt.Sprintf(":%s", env.Get("APP_PORT"))
	server := http.Server{
		Addr:    addr,
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	pkg.PanicIfError(err)
}
