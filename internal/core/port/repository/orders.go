package repository

import "go_playground/internal/core/entity"

type OrderRepository interface {
	Create(model entity.Order) error
}
