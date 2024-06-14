package server

import (
	"fmt"
	"go_playground/internal/infrastructure/config"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HttpServer interface {
	Start()
	Stop()
}
type httpServer struct {
	config *config.Config
	app    *fiber.App
}

func NewHttpServer(cfg *config.Config) (HttpServer, *fiber.App) {
	server := fiber.New(
		fiber.Config{
			AppName: cfg.App.Name,
		},
	)
	return &httpServer{
		config: cfg,
		app:    server,
	}, server
}

func (h *httpServer) Start() {

	h.app.Use(helmet.New())
	h.app.Use(recover.New())
	err := h.app.Listen(fmt.Sprintf(":%v", h.config.App.Port))
	if err != nil {
		log.Fatalf("Cannot server server ... %v", err)
	}
}

func (h *httpServer) Stop() {
	h.app.ShutdownWithTimeout(2 * time.Second)
}
