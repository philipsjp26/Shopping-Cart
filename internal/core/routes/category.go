package routes

import (
	"go_playground/internal/controller/http/products"
	"go_playground/internal/core/usecase/category"
	"go_playground/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupCategory(routes fiber.Router, db *gorm.DB) {
	api := routes.Group("/api")
	v1 := api.Group("/v1")

	categories := v1.Group("/category")

	// repository
	categoryRepo := repository.NewProductCategoryRepo(db)

	// services
	categoryServices := category.NewCateogryServices(categoryRepo)

	// controller
	categoryController := products.NewCategoryController(categoryServices)

	categories.Post("/create", categoryController.Create)
	categories.Get("", categoryController.RetrieveAll)
}
