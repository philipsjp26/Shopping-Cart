package routes

import (
	"go_playground/internal/controller/http/categories"
	"go_playground/internal/controller/http/products"
	"go_playground/internal/core/middleware"
	"go_playground/internal/core/usecase/category"
	"go_playground/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupCategory(routes fiber.Router, db *gorm.DB) {
	api := routes.Group("/api")
	v1 := api.Group("/v1")

	ct := v1.Group("/category")

	// repository
	categoryRepo := repository.NewProductCategoryRepo(db)
	productRepo := repository.NewProductRepo(db)
	accessTokenRepo := repository.NewAccessTokenRepo(db)
	// services
	categoryServices := category.NewCateogryServices(categoryRepo, productRepo)
	// controller
	productController := products.NewProductController(categoryServices)
	categoryController := categories.NewCategoryController(categoryServices)

	ct.Post("/create", middleware.Authorize(accessTokenRepo), categoryController.Create)
	ct.Get("", middleware.Authorize(accessTokenRepo), categoryController.RetrieveAll)
	ct.Get("/products/:category_id", middleware.Authorize(accessTokenRepo), categoryController.RetrieveProductsOfCategory)
	ct.Post("/products", middleware.Authorize(accessTokenRepo), productController.Create)
}
