package repository

import (
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repository.OrderRepository {
	return &orderRepo{db: db}
}

func (o *orderRepo) Create(model entity.Order) error {
	var err error
	if err = o.db.Save(&model).Error; err != nil {
		return err
	}

	return nil
}
