package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	CartRequest struct {
		CartId   int
		Products []Product `json:"products"`
	}
	Product struct {
		Id       int `json:"id"`
		Quantity int `json:"quantity"`
	}
	RemoveProductFromCart struct {
		CartId    int   `json:"cart_id"`
		ProductId []int `json:"product_ids"`
	}
)

func (p RemoveProductFromCart) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.CartId, validation.Required),
		validation.Field(&p.ProductId, validation.Required),
	)
}

func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required.Error("product id is required")),
		validation.Field(&p.Quantity, validation.Required.Error("quantity is required")),
	)
}

func (p CartRequest) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Products, validation.Required, validation.Length(1, 0),
			validation.Each(validation.By(validateProduct)),
		),
	)
}
func validateProduct(value interface{}) error {
	if product, ok := value.(Product); ok {
		return product.Validate()
	}
	return fmt.Errorf("invalid product")
}
