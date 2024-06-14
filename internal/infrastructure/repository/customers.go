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

func (c *customerRepo) FindCustomerCartProducts(customer entity.Customers) ([]entity.CustomerCartProducts, error) {
	var dest []entity.CustomerCartProducts
	cols := []string{"c2.shopping_cart_id", "p.name", "p.description", "p.quantity", "p.price", "p.created_at"}
	err := c.db.Model(&customer).Select(cols).
		Joins("LEFT JOIN shopping_carts sc on sc.customer_id = customers.id").
		Joins("LEFT JOIN carts c2 ON c2.shopping_cart_id = sc.id").
		Joins("LEFT JOIN products p ON  p.id = c2.product_id").Scan(&dest).Error
	if err != nil {
		return nil, err
	}
	return dest, nil
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
