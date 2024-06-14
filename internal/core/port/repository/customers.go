/*
	 To interact with the database, we create the secondary port component, the UserRepository interface, in
		"/internal/core/port/repository" file. this interface defines the contract for CRUD operation a new user object
*/
package repository

import (
	"context"
	"go_playground/internal/core/entity"
)

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]entity.Customers, error)
	FindOne(ctx context.Context, customer entity.Customers) (*entity.Customers, error)
	Store(ctx context.Context, customer entity.Customers) error
}
