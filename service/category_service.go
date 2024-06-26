package service

import (
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/model/dto"
)

type CategoryService interface {
	Create(ctx *fiber.Ctx, request dto.CategoryCreateRequest) dto.CategoryResponse
	Update(ctx *fiber.Ctx, request dto.CategoryUpdateRequest) dto.CategoryResponse
	Delete(ctx *fiber.Ctx, categoryId int)
	FindById(ctx *fiber.Ctx, categoryId int) dto.CategoryResponse
	FindAll(ctx *fiber.Ctx) []dto.CategoryResponse
}
