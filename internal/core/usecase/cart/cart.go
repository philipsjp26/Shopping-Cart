package cart

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/repository"
	"go_playground/internal/core/port/service"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

type cartServices struct {
	repo             repository.CartRepository
	repoProduct      repository.ProductRepository
	repoShoppingCart repository.ShoppingCartRepo
}

func NewCartService(repo repository.CartRepository, repoProduct repository.ProductRepository, shop repository.ShoppingCartRepo) service.CartService {
	return &cartServices{repo: repo, repoProduct: repoProduct, repoShoppingCart: shop}
}

func (carts *cartServices) AddProductsCart(req model.CartRequest) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)
	baseResponse.Code = http.StatusCreated
	baseResponse.Message = "success add products to cart"

	shoppingCart, err := carts.repoShoppingCart.FindOne(entity.ShoppingCart{ID: uint(req.CartId)})
	if err != nil {
		log.Errorf("find shopping cart got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	if shoppingCart == nil {
		baseResponse.Code = http.StatusNotFound
		baseResponse.Message = "cart not found"
		return &baseResponse
	}

	productId := make([]int, 0)
	for _, v := range req.Products {
		productId = append(productId, v.Id)
	}

	product, err := carts.repoProduct.FindInIds(productId...)
	if err != nil {
		log.Errorf("findOne product got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	if len(product) == 0 {
		baseResponse.Code = http.StatusNotFound
		baseResponse.Message = "data not found"
		return &baseResponse
	}

	var c []entity.Cart
	for _, cart := range req.Products {
		c = append(c, entity.Cart{
			ProductId:      cart.Id,
			ShoppingCartId: req.CartId,
			Quantity:       cart.Quantity,
		})
	}
	if err = carts.repo.BatchCreate(c); err != nil {
		log.Errorf("failed batch insert carts got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	return &baseResponse
}
