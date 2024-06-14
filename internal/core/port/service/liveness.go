package service

import "go_playground/internal/core/common"

type LivenessService interface {
	GetLiveness() *common.BaseResponse
}
