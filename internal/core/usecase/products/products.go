package products

import (
	"go_playground/internal/core/port/repository"
	"go_playground/internal/core/port/service"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductServices(repo repository.ProductRepository) service.ProductsServices {
	return &productService{repo: repo}
}
