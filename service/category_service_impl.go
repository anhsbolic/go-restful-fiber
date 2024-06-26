package service

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go-restful-fiber/model/domain"
	"go-restful-fiber/model/dto"
	"go-restful-fiber/pkg"
	"go-restful-fiber/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
	RedisClient        *redis.Client
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate, redisClient *redis.Client) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
		RedisClient:        redisClient,
	}
}

func (service *CategoryServiceImpl) Create(ctx *fiber.Ctx, request dto.CategoryCreateRequest) (dto.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return dto.CategoryResponse{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	defer pkg.CommitOrRollback(tx)
	category := domain.Category{
		Name: request.Name,
	}
	category = service.CategoryRepository.Save(ctx, tx, category)

	return dto.ToCategoryResponse(category), nil
}

func (service *CategoryServiceImpl) Update(ctx *fiber.Ctx, request dto.CategoryUpdateRequest) (dto.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return dto.CategoryResponse{}, fiber.ErrBadRequest
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	defer pkg.CommitOrRollback(tx)
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return dto.CategoryResponse{}, fiber.ErrNotFound
	}
	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)

	return dto.ToCategoryResponse(category), nil
}

func (service *CategoryServiceImpl) Delete(ctx *fiber.Ctx, categoryId int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}

	defer pkg.CommitOrRollback(tx)
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return fiber.ErrNotFound
	}
	service.CategoryRepository.Delete(ctx, tx, category)

	return nil
}

func (service *CategoryServiceImpl) FindById(ctx *fiber.Ctx, categoryId int) (dto.CategoryResponse, error) {
	// Init Category
	var category domain.Category

	// get from Redis
	redisKey := fmt.Sprintf("category:%d", categoryId)
	err := service.RedisClient.HGetAll(ctx.Context(), redisKey).Scan(&category)
	if err == nil {
		return dto.ToCategoryResponse(category), nil
	}

	// get from DB
	tx, err := service.DB.Begin()
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	defer pkg.CommitOrRollback(tx)
	category, err = service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return dto.CategoryResponse{}, fiber.ErrNotFound
	}

	// save in redis
	service.RedisClient.HSet(ctx.Context(), redisKey, category)

	// return
	pkg.NewLogger().Info("GET FROM DB")
	return dto.ToCategoryResponse(category), nil
}

func (service *CategoryServiceImpl) FindAll(ctx *fiber.Ctx) ([]dto.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []dto.CategoryResponse{}, err
	}

	defer pkg.CommitOrRollback(tx)
	categories := service.CategoryRepository.FindAll(ctx, tx)

	return dto.ToCategoryResponses(categories), nil
}
