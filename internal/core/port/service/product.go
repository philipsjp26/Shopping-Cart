package service

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
)

type ProductServices interface {
	ListProducts() *common.BaseResponse
	Create(req model.ProductCategoryRequest) *common.BaseResponse
}
