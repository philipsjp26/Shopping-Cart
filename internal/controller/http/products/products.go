package products

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	Create(c *fiber.Ctx) error
}
