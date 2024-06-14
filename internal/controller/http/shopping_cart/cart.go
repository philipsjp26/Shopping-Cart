package cart

import (
	"go_playground/internal/core/common/jwt"
	"go_playground/internal/core/port/service"

	"github.com/gofiber/fiber/v2"
)

type CartController interface {
	Create(c *fiber.Ctx) error
}
type cartController struct {
	sv service.ShoppingCartService
}

func NewCartController(service service.ShoppingCartService) CartController {
	return &cartController{sv: service}
}

func (cart *cartController) Create(c *fiber.Ctx) error {
	customer := c.Locals("customer")
	currentUser := customer.(*jwt.Customer)

	res := cart.sv.Create(currentUser.Id)
	return c.Status(res.Code).JSON(res)
}
