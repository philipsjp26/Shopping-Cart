package http

import (
	"go_playground/internal/core/port/service"

	"github.com/gofiber/fiber/v2"
)

type baseController struct {
	srv service.LivenessService
}

func NewBaseController(service service.LivenessService) *baseController {
	return &baseController{
		srv: service,
	}
}

func (c *baseController) HealthCheck(appctx *fiber.Ctx) error {
	res := c.srv.GetLiveness()
	return appctx.Status(res.Code).JSON(res)
}
