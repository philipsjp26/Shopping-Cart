package repository

import (
	"context"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) repository.CustomerRepository {
	return &customerRepo{db: db}
}

func (c *customerRepo) FindAll(ctx context.Context) ([]entity.Customers, error) {
	var (
		dest []entity.Customers
	)
	err := c.db.Model(&entity.Customers{}).Find(&dest).Error
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func (c *customerRepo) FindOne(ctx context.Context, param entity.Customers) (*entity.Customers, error) {
	var (
		dest entity.Customers
	)
	err := c.db.Where(&param).First(&dest).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &dest, nil
}

func (c *customerRepo) Store(ctx context.Context, customer entity.Customers) error {
	err := c.db.Create(&customer).Error
	if err != nil {
		return err
	}
	return nil
}
