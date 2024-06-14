package repository

import (
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"gorm.io/gorm"
)

type productCategoriesRepo struct {
	db *gorm.DB
}

func NewProductCategoryRepo(db *gorm.DB) repository.ProductCategories {
	return &productCategoriesRepo{db: db}
}
func (p *productCategoriesRepo) FindProductsCategory(id int) ([]entity.Products, error) {
	var (
		dest []entity.Products
	)
	if err := p.db.Where("category_id = ?", id).Find(&dest).Error; err != nil {
		return nil, err
	}
	return dest, nil

}
func (p *productCategoriesRepo) Store(param entity.Categories) error {
	if err := p.db.Create(&param).Error; err != nil {
		return err
	}
	return nil

}

func (p *productCategoriesRepo) FindAll(param entity.Categories) ([]entity.Categories, error) {
	var (
		dest []entity.Categories
	)
	if err := p.db.Model(&entity.Categories{}).Where(&param).Find(&dest).Error; err != nil {
		return dest, err
	}
	return dest, nil
}

func (p *productCategoriesRepo) FindOne(param entity.Categories) (*entity.Categories, error) {
	var (
		dest entity.Categories
	)
	err := p.db.Model(&entity.Categories{}).First(&dest).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &dest, nil

}
