package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(c *fiber.Ctx) error
}
