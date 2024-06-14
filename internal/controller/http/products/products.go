package products

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	Create(c *fiber.Ctx) error
}

type productController struct {
	productService service.ProductCategoryServices
}

func NewProductController(sv service.ProductCategoryServices) ProductController {
	return &productController{productService: sv}
}

func (pc *productController) Create(c *fiber.Ctx) error {
	var (
		err      error
		response common.BaseResponse
	)
	param := new(model.ProductOfCategoryRequest)
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
	res := pc.productService.CreateProductsOfCategory(param.CategoryId, *param)
	return c.Status(res.Code).JSON(res)
}
