package products_category

import "github.com/gofiber/fiber/v2"

func (cc *categoryController) RetrieveAll(c *fiber.Ctx) error {
	res := cc.categoryService.ListProducts()
	return c.Status(res.Code).JSON(res)
}
