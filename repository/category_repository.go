package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/model/domain"
)

type CategoryRepository interface {
	Save(ctx *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx *fiber.Ctx, tx *sql.Tx, category domain.Category)
	FindById(ctx *fiber.Ctx, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) []domain.Category
}
