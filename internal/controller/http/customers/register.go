package customers

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (cu *customerController) RegisterCustomer(c *fiber.Ctx) error {
	var (
		err      error
		response common.BaseResponse
	)
	param := new(model.CustomerRegisterRequest)

	if err = c.BodyParser(param); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "bad request"
		return c.Status(response.Code).JSON(response)
	}
	if err = param.Validate(); err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = common.RespValidationError
		response.Errors = err
		return c.Status(response.Code).JSON(response)
	}

	r := cu.customer.CreateCustomer(*param)
	return c.Status(r.Code).JSON(r)
}
