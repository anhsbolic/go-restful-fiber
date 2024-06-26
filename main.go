package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/config"
	"go-restful-fiber/middleware"
	"go-restful-fiber/pkg"
	"go-restful-fiber/routes"
	"time"
)

func main() {
	// Get Config
	env := config.GetEnvConfig()

	// Setup Server
	addr := fmt.Sprintf(":%s", env.Get("APP_PORT"))
	server := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Prefork:      true,
		ErrorHandler: pkg.NewErrorHandler,
	})

	// Set Global Middleware
	server.Use(middleware.XApiKeyMiddleware)

	// Setup DB
	db := pkg.NewDB()

	// Setup Redis
	redisClient := pkg.NewRedisClient()

	// Setup Validator
	validate := validator.New()

	// Setup Routes
	routes.InitCategoryRoutes(server, db, validate, redisClient)

	// Start Server
	err := server.Listen(addr)
	pkg.PanicIfError(err)
}
