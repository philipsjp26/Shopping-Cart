package cart

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func (carts *cartServices) RemoveItem(req model.RemoveProductFromCart) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)
	temp := []entity.Cart{}
	for _, v := range req.ProductId {
		temp = append(temp, entity.Cart{
			ProductId:      v,
			ShoppingCartId: req.CartId,
		})
	}
	err := carts.repo.Delete(nil, temp...)
	if err != nil {
		log.Errorf("failed delete product from cart got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = common.RespInternalServerEror
		return &baseResponse
	}
	baseResponse.Code = http.StatusOK
	baseResponse.Message = "success"
	return &baseResponse
}
