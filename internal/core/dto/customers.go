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

func EntityCustomerCartProductsToCarts(result []entity.CustomerCartProducts) model.CustomerCartProductsResponse {
	x := model.CustomerCartProductsResponse{}
	cartMap := make(map[int][]model.ProductDetail)

	for _, v := range result {
		product := model.ProductDetail{
			Name:        v.Name,
			Description: v.Description,
		}
		cartMap[v.ShoppingCartId] = append(cartMap[v.ShoppingCartId], product)
	}

	for cartID, products := range cartMap {
		x.Carts = append(x.Carts, model.Carts{
			Id:       cartID,
			Products: products,
		})
	}

	return x
}
