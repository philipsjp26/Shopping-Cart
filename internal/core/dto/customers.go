package dto

import (
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
)

func EntityToDTOCustomer(result entity.Customers) model.CustomersDetail {
	x := model.CustomersDetail{
		Id:        int(result.Id),
		Email:     result.Email,
		Phone:     *result.Phone,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
	}
	return x
}

func EntityToDTOCustomers(result []entity.Customers) []model.CustomersDetail {
	var customers []model.CustomersDetail
	for i := 0; i < len(result); i++ {
		customers = append(customers, EntityToDTOCustomer(result[i]))
	}
	return customers
}
