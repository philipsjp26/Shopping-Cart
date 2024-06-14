package repository

import "go_playground/internal/core/entity"

type ProductCategories interface {
	FindAll(param entity.Categories) ([]entity.Categories, error)
	Store(param entity.Categories) error
	FindOne(param entity.Categories) (*entity.Categories, error)
	FindProductsCategory(id int) ([]entity.Products, error)
}
