package service

import "go_playground/internal/core/common"

type ShoppingCartService interface {
	Create(custId int) *common.BaseResponse
}
