package routes

import (
	cart "go_playground/internal/controller/http/shopping_cart"
	"go_playground/internal/core/middleware"
	cartServices "go_playground/internal/core/usecase/cart"
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
	shopcartRepo := repository.NewShoppingCartRepo(db)
	accessTokenRepo := repository.NewAccessTokenRepo(db)
	cartRepo := repository.NewCartRepo(db)
	productRepo := repository.NewProductRepo(db)

	// services
	shopcartService := shoppingcart.NewShoppingCartService(shopcartRepo)
	cartServices := cartServices.NewCartService(cartRepo, productRepo, shopcartRepo)
	// controller
	cartController := cart.NewCartController(shopcartService, cartServices)

	carts.Post("", middleware.Authorize(accessTokenRepo), cartController.Create)
	carts.Post("/:cart_id", cartController.AddProducts)
}
