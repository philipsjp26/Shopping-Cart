package repository

import "go_playground/internal/core/entity"

type ShoppingCartRepo interface {
	Create(model entity.ShoppingCart) (int, error)
}
