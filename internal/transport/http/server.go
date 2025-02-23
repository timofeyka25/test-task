package http

import (
	"github.com/gofiber/fiber/v2"
	"test-task/internal/services"
	"test-task/internal/transport/http/middleware"
	"time"
)

type Server struct {
	app *fiber.App
	cfg *Config
}

func NewServer(cfg *Config, handlers []Handler, accessConfig *services.AccessConfig) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  30 * time.Second,
	})

	api := app.Group("")

	api.Use(func(c *fiber.Ctx) error {
		excludedPaths := map[string]bool{
			"/health":  true,
			"/sign-in": true,
			"/sign-up": true,
		}

		if excludedPaths[c.OriginalURL()] {
			return c.Next()
		}

		return middleware.AuthMiddleware(accessConfig.AccessSecret)(c)
	})

	for _, handler := range handlers {
		handler.Register(api)
	}
	api.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Route Not Found",
		})
	})

	return &Server{
		app: app,
		cfg: cfg,
	}
}
