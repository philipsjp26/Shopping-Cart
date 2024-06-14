package categories

import "github.com/gofiber/fiber/v2"

func (cc *categoryController) RetrieveAll(c *fiber.Ctx) error {
	res := cc.categoryService.RetrieveCategories()
	return c.Status(res.Code).JSON(res)
}
