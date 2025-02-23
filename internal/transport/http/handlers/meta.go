package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type metaHandler struct{}

func NewMetaHandler() *metaHandler {
	return &metaHandler{}
}

func (h *metaHandler) Register(router fiber.Router) {
	log.Println("Health handler")

	router.Get("/health", h.health)
}

func (h *metaHandler) health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK")
}
