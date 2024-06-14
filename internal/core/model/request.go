package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	Paging struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}
	CustomerRegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone,omitempty"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	CustomerLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	ProductCategoryRequest struct {
		Name string `json:"name"`
	}
	ProductCategoryQuery struct {
		Paging
		Id   int    `json:"id,omitempty"`
		Name string `json:"omitempty"`
	}
)

func (v ProductCategoryRequest) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Required.Error("product name is required")),
	)
}

func (v CustomerLoginRequest) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Username, validation.Required.Error("username is required")),
		validation.Field(&v.Password, validation.Required.Error("password is required")),
	)
}

func (v CustomerRegisterRequest) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required.Error("email is required")),
		validation.Field(&v.Name, validation.Required.Error("name is required")),
		validation.Field(&v.Username, validation.Required.Error("username is required")),
		validation.Field(&v.Password, validation.Required.Error("password is required")),
	)
}
