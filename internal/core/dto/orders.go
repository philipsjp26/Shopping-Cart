package dto

import "go_playground/internal/core/model"

func DTOOrderItemsResponse(trx, status string) model.OrderItemsResponse {
	return model.OrderItemsResponse{
		TransactionId: trx,
		Status:        status,
	}
}
