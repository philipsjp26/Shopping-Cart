package cart

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/common/jwt"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CartController interface {
	Create(c *fiber.Ctx) error
	AddProducts(c *fiber.Ctx) error
}
type cartController struct {
	sv     service.ShoppingCartService
	cartSv service.CartService
}

func NewCartController(service service.ShoppingCartService, cartsv service.CartService) CartController {
	return &cartController{sv: service, cartSv: cartsv}
}

func (cart *cartController) Create(c *fiber.Ctx) error {
	customer := c.Locals("customer")
	currentUser := customer.(*jwt.Customer)

	res := cart.sv.Create(currentUser.Id)
	return c.Status(res.Code).JSON(res)
}

func (cart *cartController) AddProducts(c *fiber.Ctx) error {
	var (
		err      error
		response common.BaseResponse
		cartId   = c.Params("cart_id")
	)
	param := new(model.CartRequest)

	if err = c.BodyParser(param); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "bad request"
		return c.Status(response.Code).JSON(response)
	}
	if err = param.Validate(); err != nil || cartId == "" {
		response.Code = http.StatusUnprocessableEntity
		response.Message = common.RespValidationError
		response.Errors = err
		return c.Status(response.Code).JSON(response)
	}

	id, _ := strconv.Atoi(cartId)
	param.CartId = id
	res := cart.cartSv.AddProductsCart(*param)
	return c.Status(res.Code).JSON(res)
}
