package repository

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) []domain.Category
}
