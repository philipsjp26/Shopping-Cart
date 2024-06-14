package service

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
)

type (
	ProductsServices interface {
		Create(req model.ProductRequest) *common.BaseResponse
	}
)
