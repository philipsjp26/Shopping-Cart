package repository

import "go_playground/internal/core/entity"

type CartRepository interface {
	Create(param entity.Cart) error
	BatchCreate(param []entity.Cart) error
	Delete(param *entity.Cart, bulk ...entity.Cart) error
	FindProductsOfCart(param entity.Cart) ([]entity.ProductsOfCart, error)
}
