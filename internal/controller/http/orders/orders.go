package orders

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/common/jwt"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	Create(c *fiber.Ctx) error
}

type orderController struct {
	orderService service.OrderServices
}

func NewOrderController(sv service.OrderServices) OrderController {
	return &orderController{orderService: sv}
}

func (o *orderController) Create(c *fiber.Ctx) error {
	var (
		err      error
		response common.BaseResponse
	)
	currentUser := c.Locals("customer")
	customer := currentUser.(*jwt.Customer)
	param := new(model.OrderRequest)
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
	param.CustomerId = customer.Id
	res := o.orderService.CreateOrders(*param)
	return c.Status(res.Code).JSON(res)
}
