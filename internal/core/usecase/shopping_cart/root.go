package shoppingcart

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"
	"go_playground/internal/core/port/service"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

type shoppingCartService struct {
	repo repository.ShoppingCartRepo
}

func NewShoppingCartService(repo repository.ShoppingCartRepo) service.ShoppingCartService {
	return &shoppingCartService{repo: repo}
}

func (scs *shoppingCartService) Create(custId int) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)
	id, err := scs.repo.Create(entity.ShoppingCart{CustomerId: custId})
	if err != nil {
		log.Errorf("failed store shopping cart got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = common.RespInternalServerEror
		return &baseResponse
	}
	baseResponse.Code = http.StatusCreated
	baseResponse.Message = "success create cart"
	baseResponse.Data = map[string]any{
		"cart_id": id,
	}
	return &baseResponse
}
