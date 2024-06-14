package category

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/repository"
	"go_playground/internal/core/port/service"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

type categoryServices struct {
	repo repository.ProductCategories
}

func NewCateogryServices(r repository.ProductCategories) service.ProductServices {
	return &categoryServices{repo: r}
}

func (cs *categoryServices) Create(req model.ProductCategoryRequest) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)
	category, err := cs.repo.FindOne(entity.Categories{
		Name: req.Name,
	})
	if err != nil {
		log.Errorf("find repo product categories got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	if category != nil {
		baseResponse.Code = http.StatusConflict
		baseResponse.Message = "category already exists"
		return &baseResponse
	}
	if err = cs.repo.Store(entity.Categories{
		Name: req.Name,
	}); err != nil {
		log.Errorf("failed store category got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	baseResponse.Code = http.StatusCreated
	baseResponse.Message = "success create category"
	return &baseResponse
}

func (cs *categoryServices) ListProducts() *common.BaseResponse {
	var (
		response common.BaseResponse
	)
	categories, err := cs.repo.FindAll(entity.Categories{})
	if err != nil {
		log.Error(err)
		response.Code = http.StatusInternalServerError
		response.Message = "failed retrieve categories"
		return &response

	}
	response.Code = http.StatusOK
	response.Message = "success retrieve categories"
	response.Data = categories
	return &response
}
