package model

import validation "github.com/go-ozzo/ozzo-validation"

func (v ProductRequest) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.CategoryId, validation.Required.Error("Category ID is required")),
		validation.Field(&v.Name, validation.Required.Error("Name is required")),
		validation.Field(&v.Price, validation.Required.Error("Price is required")),
		validation.Field(&v.Quantity, validation.Required.Error("Quantity is required")),
	)
}
