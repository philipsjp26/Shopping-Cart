package http

import (
	"go_playground/internal/core/server"

	"go_playground/internal/infrastructure/config"
	"go_playground/internal/infrastructure/repository"
)

func Runner() {
	cfg := config.Configuration()

	app, _ := server.NewHttpServer(cfg)

	// Initialize db sql connection
	initDBConn := repository.NewDBGetConnection(cfg.Database.Driver)
	_ = initDBConn.Connect(cfg)

	// /* Create the service from usecase */
	// livenessPort := usecase.NewLivenessService()

	// /* Mount primary adapter */
	// routes := http.NewBaseController(server, livenessPort)
	// routes.InitRouter()

	app.Start()
}
