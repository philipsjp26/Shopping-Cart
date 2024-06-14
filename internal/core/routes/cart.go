package routes

import (
	cart "go_playground/internal/controller/http/shopping_cart"
	"go_playground/internal/core/middleware"
	shoppingcart "go_playground/internal/core/usecase/shopping_cart"
	"go_playground/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupCart(routes fiber.Router, db *gorm.DB) {
	api := routes.Group("/api")
	v1 := api.Group("/v1")
	carts := v1.Group("/carts")

	// repository
	cartRepo := repository.NewShoppingCartRepo(db)
	accessTokenRepo := repository.NewAccessTokenRepo(db)

	// services
	cartService := shoppingcart.NewShoppingCartService(cartRepo)
	// controller
	cartController := cart.NewCartController(cartService)

	carts.Post("", middleware.Authorize(accessTokenRepo), cartController.Create)
}
