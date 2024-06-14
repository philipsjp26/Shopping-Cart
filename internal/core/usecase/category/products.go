package category

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func (cs *categoryServices) RetrieveProducts(categoryId int) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)

	category, err := cs.repo.FindOne(entity.Categories{Id: uint(categoryId)})
	if err != nil {
		log.Errorf("find repo categories got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	if category == nil {
		log.Warn("category not found")
		baseResponse.Code = http.StatusNotFound
		baseResponse.Message = "category not found"
		return &baseResponse
	}
	categoryProducts, err := cs.repo.FindProductsCategory(int(category.Id))
	if err != nil {
		log.Errorf("find products of category got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	baseResponse.Code = http.StatusOK
	baseResponse.Message = "success retrieve products"
	baseResponse.Data = categoryProducts
	log.Info("success retrieve products")
	return &baseResponse
}
func (cs *categoryServices) CreateProductsOfCategory(categoryId int, req model.ProductOfCategoryRequest) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)
	if err := cs.repoProduct.Store(entity.Products{
		CategoryID:  categoryId,
		Name:        req.Name,
		Description: &req.Description,
		Quantity:    req.Quantity,
		Price:       req.Price,
	}); err != nil {
		log.Errorf("store products got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	baseResponse.Code = http.StatusCreated
	baseResponse.Message = "success create products"
	log.Info("success create products")
	return &baseResponse
}
