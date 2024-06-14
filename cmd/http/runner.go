package http

import (
	"go_playground/internal/core/routes"
	"go_playground/internal/core/server"

	"go_playground/internal/infrastructure/config"
	"go_playground/internal/infrastructure/database"
)

func Runner() {
	cfg := config.Configuration()

	app, server := server.NewHttpServer(cfg)

	// Initialize db sql connection
	initDBConn := database.NewDBGetConnection(cfg.Database.Driver)
	db := initDBConn.Connect(cfg)

	routes.SetupCustomerRoutes(server, db)
	routes.SetupCategory(server, db)
	app.Start()
}
