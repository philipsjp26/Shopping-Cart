package customers

import (
	"context"
	"go_playground/internal/core/common"
	"go_playground/internal/core/common/utils"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func (c *customerService) CreateCustomer(param model.CustomerRegisterRequest) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)

	customer, err := c.customerRepo.FindOne(context.Background(), entity.Customers{
		Email:    param.Email,
		Username: param.Username,
	})
	if err != nil {
		log.Errorf("failed find customer got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}

	if customer != nil {
		log.Warnf("customer already exists ID :%v", customer.Id)
		baseResponse.Code = http.StatusConflict
		baseResponse.Message = "customer already exists"
		return &baseResponse
	}

	param.Password = utils.BcryptGen(param.Password)
	if err = c.customerRepo.Store(context.Background(), entity.Customers{
		Email:    param.Email,
		Phone:    &param.Phone,
		Name:     param.Name,
		Username: param.Username,
		Password: param.Password,
	}); err != nil {
		log.Errorf("failed store customer got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}

	log.Info("success register customer")

	baseResponse.Code = http.StatusCreated
	baseResponse.Message = "success register customer"
	return &baseResponse
}
