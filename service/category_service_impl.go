package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-restful-fiber/exception"
	"go-restful-fiber/model/domain"
	"go-restful-fiber/model/dto"
	"go-restful-fiber/pkg"
	"go-restful-fiber/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request dto.CategoryCreateRequest) dto.CategoryResponse {
	err := service.Validate.Struct(request)
	pkg.PanicIfError(err)

	tx, err := service.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return dto.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request dto.CategoryUpdateRequest) dto.CategoryResponse {
	err := service.Validate.Struct(request)
	pkg.PanicIfError(err)

	tx, err := service.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return dto.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) dto.CategoryResponse {
	tx, err := service.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return dto.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []dto.CategoryResponse {
	tx, err := service.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return dto.ToCategoryResponses(categories)
}
