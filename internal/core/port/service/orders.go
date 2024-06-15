package service

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
)

type OrderServices interface {
	CreateOrders(request model.OrderRequest) *common.BaseResponse
}
