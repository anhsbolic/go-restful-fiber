package routes

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/controller"
	"go-restful-fiber/repository"
	"go-restful-fiber/service"
)

func InitCategoryRoutes(server *fiber.App, db *sql.DB, validate *validator.Validate) {
	// Setup Category API
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Set Routes
	v1CategoriesAPI := server.Group("/api/v1/categories")
	v1CategoriesAPI.Post("/", categoryController.Create)
	v1CategoriesAPI.Get("/", categoryController.FindAll)
	v1CategoriesAPI.Get("/:categoryId", categoryController.FindById)
}
