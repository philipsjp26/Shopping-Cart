package model

import validation "github.com/go-ozzo/ozzo-validation"

type (
	ProductOfCategoryRequest struct {
		CategoryId  int     `json:"category_id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Quantity    int     `json:"quantity"`
		Price       float64 `json:"price"`
	}
)

func (v ProductOfCategoryRequest) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.CategoryId, validation.Required.Error("category id is required")),
		validation.Field(&v.Name, validation.Required.Error("name is required")),
		validation.Field(&v.Description, validation.Required.Error("description is required")),
		validation.Field(&v.Quantity, validation.Required.Error("quantity is required")),
		validation.Field(&v.Price, validation.Required.Error("price is required")),
	)
}
