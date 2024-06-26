package repository

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/model/domain"
	"go-restful-fiber/pkg"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into categories (name) values ($1) returning id"
	err := tx.QueryRowContext(ctx.Context(), SQL, category.Name).Scan(&category.ID)
	pkg.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update categories set name = $1 where id = $2"
	_, err := tx.ExecContext(ctx.Context(), SQL, category.Name, category.ID)
	pkg.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx *fiber.Ctx, tx *sql.Tx, category domain.Category) {
	SQL := "delete from categories where id = $1"
	_, err := tx.ExecContext(ctx.Context(), SQL, category.ID)
	pkg.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx *fiber.Ctx, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from categories where id = $1"
	rows, err := tx.QueryContext(ctx.Context(), SQL, categoryId)
	pkg.PanicIfError(err)
	defer func(rows *sql.Rows) {
		if err != nil {
			pkg.PanicIfError(err)
		}
	}(rows)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		pkg.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from categories"
	rows, err := tx.QueryContext(ctx.Context(), SQL)
	pkg.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			pkg.PanicIfError(err)
		}
	}(rows)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		pkg.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
