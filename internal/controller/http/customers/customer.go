package customers

import (
	"go_playground/internal/core/port/service"

	"github.com/gofiber/fiber/v2"
)

type CustomerController interface {
	RetrieveAllCustomers(c *fiber.Ctx) error
	RegisterCustomer(c *fiber.Ctx) error
	LoginCustomer(ct *fiber.Ctx) error
	CartProducts(ct *fiber.Ctx) error
}

type customerController struct {
	customer service.CustomerServices
}

func NewCustomerController(sv service.CustomerServices) CustomerController {
	return &customerController{customer: sv}
}
