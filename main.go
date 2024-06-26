package main

import (
	"go-restful-fiber/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	//validatorOption := validator.Option()
	//validate := validator.New()
	//server := app.InitializeServer()
	//err := server.ListenAndServe()
	//helper.PanicIfError(err)
}
