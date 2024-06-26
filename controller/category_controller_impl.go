package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julienschmidt/httprouter"
	"go-restful-fiber/model/dto"
	"go-restful-fiber/pkg"
	"go-restful-fiber/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := dto.CategoryCreateRequest{}
	pkg.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	pkg.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := dto.CategoryUpdateRequest{}
	pkg.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	pkg.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	pkg.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	pkg.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := dto.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	pkg.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	pkg.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	pkg.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	categoryResponses := controller.CategoryService.FindAll(ctx)
	webResponse := dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	return ctx.JSON(webResponse)
}
