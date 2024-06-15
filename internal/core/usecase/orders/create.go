package orders

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/dto"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"go_playground/internal/core/port/repository"
	"go_playground/internal/core/port/service"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type orderServices struct {
	repo     repository.OrderRepository
	cartRepo repository.CartRepository
}

func NewOrderService(repo repository.OrderRepository, cartRepo repository.CartRepository) service.OrderServices {
	return &orderServices{repo: repo, cartRepo: cartRepo}
}

func (o *orderServices) CreateOrders(request model.OrderRequest) *common.BaseResponse {
	var (
		baseResponse common.BaseResponse
	)

	products, err := o.cartRepo.FindProductsOfCart(entity.Cart{Id: request.CartId})
	if err != nil {
		log.Errorf("failed findProductsOfCart got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = common.RespInternalServerEror
		return &baseResponse
	}
	if len(products) == 0 {
		log.Warn("products not found on cart")
		baseResponse.Code = http.StatusNotFound
		baseResponse.Message = "product not found"
		return &baseResponse
	}
	transactionNumber := uuid.NewString()
	entities := entity.Order{
		TransactionId: transactionNumber,
		CustomerId:    request.CustomerId,
		Status:        common.Pending,
		PaymentMethod: request.PaymentMethod,
		OrderDate:     time.Now(),
	}
	items := []entity.OrderItem{}

	for _, p := range products {
		items = append(items, entity.OrderItem{
			ProductId: p.ProductId,
			Quantity:  p.TotalQty,
			Price:     p.Price,
		})
	}
	entities.OrderItems = items
	log.Info(entities)
	if err = o.repo.Create(entities); err != nil {
		log.Errorf("failed store orders")
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = common.RespInternalServerEror
		return &baseResponse
	}
	baseResponse.Code = http.StatusCreated
	baseResponse.Message = "success create orders"
	baseResponse.Data = dto.DTOOrderItemsResponse(transactionNumber, common.Pending)
	return &baseResponse
}
