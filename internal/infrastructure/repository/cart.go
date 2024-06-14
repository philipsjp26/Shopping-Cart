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

func (cart *cartRepo) Delete(param *entity.Cart, bulk ...entity.Cart) error {
	var (
		err error
	)
	if len(bulk) > 0 {
		for _, cartItem := range bulk {
			err := cart.db.Where("product_id = ? AND shopping_cart_id = ?", cartItem.ProductId, cartItem.ShoppingCartId).Delete(&entity.Cart{}).Error
			if err != nil {
				return err
			}
		}
	} else {
		err = cart.db.Delete(&param).Error
	}
	if err != nil {
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
