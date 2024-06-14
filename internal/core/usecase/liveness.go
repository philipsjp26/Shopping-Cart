// Package service provides the implementation of the LivenessService interface.
// The LivenessService handles the core business logic related to system liveness.

package usecase

import (
	"go_playground/internal/core/common"
	"net/http"

	"go_playground/internal/core/port/service"
)

type livenesService struct {
}

func NewLivenessService() service.LivenessService {
	return &livenesService{}
}

func (s *livenesService) GetLiveness() *common.BaseResponse {
	var (
		res common.BaseResponse
	)
	res.Code = http.StatusOK
	res.Message = "healthy .."
	return &res
}
