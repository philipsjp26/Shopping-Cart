package categories

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductCategoryController interface {
	RetrieveAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	RetrieveProductsOfCategory(c *fiber.Ctx) error
}

type categoryController struct {
	categoryService service.ProductCategoryServices
}

func NewCategoryController(sv service.ProductCategoryServices) ProductCategoryController {
	return &categoryController{categoryService: sv}
}

func (cc *categoryController) Create(c *fiber.Ctx) error {
	var (
		err      error
		response common.BaseResponse
	)
	param := new(model.CategoryRequest)
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
	r := cc.categoryService.Create(*param)
	return c.Status(r.Code).JSON(r)
}
func (cc *categoryController) RetrieveProductsOfCategory(c *fiber.Ctx) error {
	var (
		response common.BaseResponse
	)
	id := c.Params("category_id")
	if id == "" {
		response.Code = http.StatusUnprocessableEntity
		response.Message = "id is empty"
		return c.Status(response.Code).JSON(response)
	}
	productId, _ := strconv.Atoi(id)
	result := cc.categoryService.RetrieveProducts(productId)
	return c.Status(result.Code).JSON(result)
}
