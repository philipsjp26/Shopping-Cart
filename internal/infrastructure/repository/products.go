package repository

import (
	"fmt"
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

func (pt *productRepositories) FindInIds(ids ...int) ([]entity.Products, error) {
	var dest []entity.Products

	// Jika ada IDs yang diberikan
	if len(ids) > 0 {
		err := pt.db.Select("id", "category_id", "name", "quantity", "price", "created_at", "updated_at").Where("id IN ?", ids).Find(&dest).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
		return dest, nil
	}

	// Jika tidak ada IDs yang diberikan, kembalikan error atau handle kasus ini sesuai kebutuhan
	return nil, fmt.Errorf("no ids provided")
}
