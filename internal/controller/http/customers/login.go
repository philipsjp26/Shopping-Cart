package customers

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (cu *customerController) LoginCustomer(ct *fiber.Ctx) error {
	var (
		err      error
		response common.BaseResponse
	)

	param := new(model.CustomerLoginRequest)
	if err = ct.BodyParser(param); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "bad request"
		return ct.Status(response.Code).JSON(response)
	}
	if err = param.Validate(); err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = common.RespValidationError
		response.Errors = err
		return ct.Status(response.Code).JSON(response)
	}

	res := cu.customer.Login(*param)
	return ct.Status(res.Code).JSON(res)
}
