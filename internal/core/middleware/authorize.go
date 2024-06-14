package middleware

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/common/jwt"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Authorize(repo repository.AccessTokenRepository) fiber.Handler {
	return func(req *fiber.Ctx) error {
		var (
			response common.BaseResponse
		)
		log.Info("Authorize")

		h := req.GetReqHeaders()["Authorization"]

		if len(h) == 0 {
			response.Code = http.StatusUnauthorized
			response.Message = "unauthorize"
			return req.Status(http.StatusUnauthorized).JSON(response)
		}
		token := h[0]
		bearer := strings.Split(token, " ")[1]

		if bearer == "" {
			response.Code = http.StatusUnauthorized
			response.Message = "unauthorize"
			return req.Status(http.StatusUnauthorized).JSON(response)
		}

		currentToken, err := repo.FindOne(entity.AccessToken{AccessToken: bearer})
		if err != nil {
			log.Errorf("failed find one access token got err :%v", err)
			response.Code = http.StatusUnauthorized
			response.Message = "unauthorize"
			return req.Status(http.StatusUnauthorized).JSON(response)
		}

		if currentToken == nil || currentToken.AccessToken != bearer {
			response.Code = http.StatusUnauthorized
			response.Message = "unauthorize"
			return req.Status(http.StatusUnauthorized).JSON(response)
		}

		claims, err := jwt.ParseAccessToken(bearer)
		if err != nil {
			log.Error(err)
			response.Code = http.StatusNotFound
			response.Message = "invalid credentials"
			return req.Status(response.Code).JSON(response)
		}

		req.Locals("customer", claims)
		req.Next()
		return nil
	}
}
