package products

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductCategoryController interface {
	RetrieveAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type categoryController struct {
	categoryService service.ProductServices
}

func NewCategoryController(sv service.ProductServices) ProductCategoryController {
	return &categoryController{categoryService: sv}
}

func (cc *categoryController) Create(c *fiber.Ctx) error {
	var (
		err      error
		response common.BaseResponse
	)
	param := new(model.ProductCategoryRequest)
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
