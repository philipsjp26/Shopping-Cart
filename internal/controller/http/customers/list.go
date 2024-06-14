package customers

import (
	"github.com/gofiber/fiber/v2"
)

func (cu *customerController) RetrieveAllCustomers(c *fiber.Ctx) error {
	res := cu.customer.GetAllCustomer()
	return c.Status(res.Code).JSON(res)
}
