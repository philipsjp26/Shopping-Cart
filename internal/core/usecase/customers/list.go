package customers

import (
	"context"
	"go_playground/internal/core/common"
	"go_playground/internal/core/dto"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func (c *customerService) GetAllCustomer() *common.BaseResponse {
	var (
		response common.BaseResponse
	)
	customers, err := c.customerRepo.FindAll(context.Background())
	if err != nil {
		log.Error(err)
		response.Code = http.StatusInternalServerError
		response.Message = "failed retrieve customers"
		return &response

	}
	response.Code = http.StatusOK
	response.Message = "success retrieve customers"
	response.Data = dto.EntityToDTOCustomers(customers)
	return &response
}
