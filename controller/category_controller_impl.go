package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/model/dto"
	"go-restful-fiber/pkg"
	"go-restful-fiber/service"
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

func (controller *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	panic("implement me")
	//categoryCreateRequest := dto.CategoryCreateRequest{}
	//pkg.ReadFromRequestBody(request, &categoryCreateRequest)
	//
	//categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	//webResponse := dto.ApiResponse{
	//	Code:   200,
	//	Status: "OK",
	//	Data:   categoryResponse,
	//}
	//
	//pkg.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	panic("implement me")
	//categoryUpdateRequest := dto.CategoryUpdateRequest{}
	//pkg.ReadFromRequestBody(request, &categoryUpdateRequest)
	//
	//categoryId := params.ByName("categoryId")
	//id, err := strconv.Atoi(categoryId)
	//pkg.PanicIfError(err)
	//
	//categoryUpdateRequest.Id = id
	//
	//categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	//webResponse := dto.ApiResponse{
	//	Code:   200,
	//	Status: "OK",
	//	Data:   categoryResponse,
	//}
	//
	//pkg.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	panic("implement me")
	//categoryId := params.ByName("categoryId")
	//id, err := strconv.Atoi(categoryId)
	//pkg.PanicIfError(err)
	//
	//controller.CategoryService.Delete(request.Context(), id)
	//webResponse := dto.ApiResponse{
	//	Code:   200,
	//	Status: "OK",
	//}
	//
	//pkg.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")
	id, err := strconv.Atoi(categoryId)
	pkg.PanicIfError(err)

	result := controller.CategoryService.FindById(ctx, id)
	return ctx.JSON(dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}

func (controller *CategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	result := controller.CategoryService.FindAll(ctx)
	return ctx.JSON(dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}
