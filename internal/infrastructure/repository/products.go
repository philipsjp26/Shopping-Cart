package repository

import (
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"gorm.io/gorm"
)

type productRepositories struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repository.ProductRepository {
	return &productRepositories{db: db}
}

func (pt *productRepositories) Store(param entity.Products) error {
	if err := pt.db.Create(&param).Error; err != nil {
		return err
	}
	return nil
}
