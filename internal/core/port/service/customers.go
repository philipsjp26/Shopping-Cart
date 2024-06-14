package service

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
)

type CustomerServices interface {
	GetAllCustomer() *common.BaseResponse
	CreateCustomer(param model.CustomerRegisterRequest) *common.BaseResponse
	Login(param model.CustomerLoginRequest) *common.BaseResponse
}
