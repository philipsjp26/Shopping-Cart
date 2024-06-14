package service

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
)

type CartService interface {
	AddProductsCart(req model.CartRequest) *common.BaseResponse
	RemoveItem(req model.RemoveProductFromCart) *common.BaseResponse
}
