package customers

import (
	"go_playground/internal/core/common/jwt"

	"github.com/gofiber/fiber/v2"
)

func (cu *customerController) CartProducts(ct *fiber.Ctx) error {
	currentUser := ct.Locals("customer")
	customer := currentUser.(*jwt.Customer)

	res := cu.customer.GetCartOfProducts(customer.Id)
	return ct.Status(res.Code).JSON(res)

}
