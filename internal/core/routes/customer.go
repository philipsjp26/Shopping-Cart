package routes

import (
	"go_playground/internal/controller/http"
	customerController "go_playground/internal/controller/http/customers"
	"go_playground/internal/core/middleware"
	"go_playground/internal/core/usecase"
	"go_playground/internal/core/usecase/customers"
	"go_playground/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupCustomerRoutes(routes fiber.Router, db *gorm.DB) {
	api := routes.Group("/api")
	v1 := api.Group("/v1")

	// repository
	customerRepo := repository.NewCustomerRepo(db)
	accessTokenRepo := repository.NewAccessTokenRepo(db)

	// end of line repository

	// services
	customerService := customers.NewCustomerServices(customerRepo, accessTokenRepo)
	livenessService := usecase.NewLivenessService()
	// end of line services

	// controller
	c := customerController.NewCustomerController(customerService)
	liveness := http.NewBaseController(livenessService)
	// end of line controller

	// health check
	api.Get("/health", liveness.HealthCheck)

	// routes

	v1Cust := v1.Group("/customers")
	v1Cust.Get("", middleware.Authorize(accessTokenRepo), c.RetrieveAllCustomers)

	v1Cust.Post("/register", c.RegisterCustomer)
	v1Cust.Post("/login", c.LoginCustomer)
	v1Cust.Get("/carts", middleware.Authorize(accessTokenRepo), c.CartProducts)

}
