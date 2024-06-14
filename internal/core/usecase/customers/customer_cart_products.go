package customers

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/dto"
	"go_playground/internal/core/entity"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func (c *customerService) GetCartOfProducts(custId int) *common.BaseResponse {
	var (
		response common.BaseResponse
	)
	carts, err := c.customerRepo.FindCustomerCartProducts(entity.Customers{Id: uint(custId)})
	if err != nil {
		log.Error(err)
		response.Code = http.StatusInternalServerError
		response.Message = "failed cart of products"
		return &response
	}

	response.Code = http.StatusOK
	response.Message = "success retrieve data"
	response.Data = dto.EntityCustomerCartProductsToCarts(carts)
	return &response
}
