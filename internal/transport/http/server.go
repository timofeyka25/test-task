package http

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Server struct {
	app *fiber.App
	cfg *Config
}

func NewServer(cfg *Config, handlers []Handler) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  30 * time.Second,
	})

	api := app.Group("")

	for _, handler := range handlers {
		handler.Register(api)
	}

	return &Server{
		app: app,
		cfg: cfg,
	}
}
