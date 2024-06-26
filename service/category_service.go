package service

import (
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/model/dto"
)

type CategoryService interface {
	Create(ctx *fiber.Ctx, request dto.CategoryCreateRequest) (dto.CategoryResponse, error)
	Update(ctx *fiber.Ctx, request dto.CategoryUpdateRequest) (dto.CategoryResponse, error)
	Delete(ctx *fiber.Ctx, categoryId int) error
	FindById(ctx *fiber.Ctx, categoryId int) (dto.CategoryResponse, error)
	FindAll(ctx *fiber.Ctx) ([]dto.CategoryResponse, error)
}
