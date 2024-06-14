package customers

import (
	"go_playground/internal/core/port/repository"
	"go_playground/internal/core/port/service"
)

type customerService struct {
	customerRepo    repository.CustomerRepository
	accessTokenRepo repository.AccessTokenRepository
}

func NewCustomerServices(repo repository.CustomerRepository, accessTokenRepo repository.AccessTokenRepository) service.CustomerServices {

	return &customerService{
		customerRepo:    repo,
		accessTokenRepo: accessTokenRepo,
	}
}
