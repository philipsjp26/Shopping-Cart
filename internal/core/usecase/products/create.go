package products

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func (ps *productService) Create(req model.ProductRequest) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)
	err := ps.repo.Store(entity.Products{
		CategoryID:  req.CategoryId,
		Name:        req.Name,
		Description: &req.Description,
		Quantity:    req.Quantity,
		Price:       req.Price,
	})
	if err != nil {
		log.Errorf("failed store products got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	baseResponse.Code = http.StatusCreated
	baseResponse.Message = "success create products"
	return &baseResponse
}
