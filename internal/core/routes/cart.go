package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupCart(routes fiber.Router, db *gorm.DB) {
	api := routes.Group("/api")
	v1 := api.Group("/v1")
	_ = v1.Group("/carts")

}
