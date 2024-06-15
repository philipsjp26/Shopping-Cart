package service

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
)

type ProductCategoryServices interface {
	RetrieveCategories() *common.BaseResponse
	Create(req model.CategoryRequest) *common.BaseResponse
	RetrieveProducts(categoryId int) *common.BaseResponse
	CreateProductsOfCategory(categoryId int, req model.ProductOfCategoryRequest) *common.BaseResponse
}
