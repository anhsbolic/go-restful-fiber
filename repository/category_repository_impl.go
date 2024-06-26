package repository

import (
	"context"
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

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into categories (name) values ($1) returning id"
	err := tx.QueryRowContext(ctx, SQL, category.Name).Scan(&category.ID)
	pkg.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)
	pkg.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.ID)
	pkg.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	pkg.PanicIfError(err)
	defer rows.Close()

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
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		pkg.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
