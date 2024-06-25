package main

import (
	"go-restful-fiber/helper"
	"net/http"
)

func main() {
	// Setup DB
	//db := app.NewDB()

	// Setup Validator
	//validate := validator.New()

	// Setup Category API
	//categoryRepository := repository.NewCategoryRepository()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//router := app.NewRouter(categoryController)

	// Setup Server
	server := http.Server{
		Addr: "localhost:3000",
		//Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
