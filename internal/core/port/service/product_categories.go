package service

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
)

type ProductCategoryServices interface {
	ListProducts() *common.BaseResponse
	Create(req model.ProductCategoryRequest) *common.BaseResponse
}
