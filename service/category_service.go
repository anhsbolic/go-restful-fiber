package service

import (
	"context"
	"go-restful-fiber/model/dto"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CategoryCreateRequest) dto.CategoryResponse
	Update(ctx context.Context, request dto.CategoryUpdateRequest) dto.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) dto.CategoryResponse
	FindAll(ctx context.Context) []dto.CategoryResponse
}
