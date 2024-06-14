package repository

import (
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"gorm.io/gorm"
)

type shoppingCartRepo struct {
	db *gorm.DB
}

func NewShoppingCartRepo(db *gorm.DB) repository.ShoppingCartRepo {
	return &shoppingCartRepo{db: db}
}

func (sc *shoppingCartRepo) Create(model entity.ShoppingCart) (int, error) {
	if err := sc.db.Create(&model).Error; err != nil {
		return 0, err
	}

	return int(model.ID), nil
}

func (sc *shoppingCartRepo) FindOne(model entity.ShoppingCart) (*entity.ShoppingCart, error) {
	var (
		dest entity.ShoppingCart
	)
	err := sc.db.Where(&model).First(&dest).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &dest, err
}
