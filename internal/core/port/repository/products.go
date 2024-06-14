package repository

import "go_playground/internal/core/entity"

type (
	ProductRepository interface {
		Store(param entity.Products) error
	}
)
