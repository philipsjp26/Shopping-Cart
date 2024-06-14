package customers

import (
	"context"
	"go_playground/internal/core/common"
	"go_playground/internal/core/common/jwt"
	"go_playground/internal/core/common/utils"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

func (c *customerService) Login(param model.CustomerLoginRequest) *common.BaseResponse {
	var (
		baseResponse              common.BaseResponse
		wg                        sync.WaitGroup
		errChan                   = make(chan error, 2)
		accessToken, refreshToken string
	)
	customer, err := c.customerRepo.FindOne(context.Background(), entity.Customers{
		Username: param.Username,
	})
	if err != nil {
		log.Errorf("error find customer got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse

	}
	if customer == nil {
		log.Warn("customer not found")
		baseResponse.Code = http.StatusNotFound
		baseResponse.Message = "data not found"
		return &baseResponse
	}
	if err = utils.BcryptVerify(param.Password, customer.Password); err != nil {
		log.Errorf("invalid password :%v", err)
		baseResponse.Code = http.StatusUnprocessableEntity
		baseResponse.Message = "invalid password"
		return &baseResponse
	}

	wg.Add(2)

	go func(wg *sync.WaitGroup, channel chan<- error) {
		defer wg.Done()
		accessToken, err = jwt.GenerateAccessToken(*customer)
		if err != nil {
			log.Errorf("failed generate access token got err :%v", err)
			errChan <- err
		}

	}(&wg, errChan)

	go func(wg *sync.WaitGroup, channel chan<- error) {
		defer wg.Done()
		refreshToken, err = jwt.GenerateRefreshToken()
		if err != nil {
			log.Errorf("failed generate refresh token got err :%v", err)
			errChan <- err
		}
	}(&wg, errChan)

	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			log.Error(err)
			baseResponse.Code = http.StatusInternalServerError
			baseResponse.Message = "internal server error"
			return &baseResponse
		}
	}
	if err = c.accessTokenRepo.Store(context.Background(), entity.AccessToken{
		CustomerID:   customer.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}); err != nil {
		log.Errorf("failed store access token got err :%v", err)
		baseResponse.Code = http.StatusInternalServerError
		baseResponse.Message = "internal server error"
		return &baseResponse
	}
	log.Info("success create login")

	baseResponse.Code = http.StatusOK
	baseResponse.Message = "success"
	baseResponse.Data = model.CustomerLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    "7200",
	}
	return &baseResponse
}
