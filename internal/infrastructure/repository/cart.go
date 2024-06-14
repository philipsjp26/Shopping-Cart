package repository

import (
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) repository.CartRepository {
	return &cartRepo{db: db}
}
func (cart *cartRepo) Create(param entity.Cart) error {
	if err := cart.db.Create(&param).Error; err != nil {
		return err
	}
	return nil

}

func (cart *cartRepo) BatchCreate(param []entity.Cart) error {
	tx := cart.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Melakukan operasi create
	if err := tx.Create(&param).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaksi jika tidak ada kesalahan
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
