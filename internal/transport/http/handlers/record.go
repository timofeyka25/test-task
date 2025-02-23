package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"test-task/internal/services"
)

type recordHandler struct {
	recordService *services.RecordService
}

func NewRecordHandler(recordService *services.RecordService) *recordHandler {
	return &recordHandler{recordService: recordService}
}

func (h *recordHandler) Register(router fiber.Router) {
	records := router.Group("/records")
	records.Get("/all", h.getAll)
}

func (h *recordHandler) getAll(ctx *fiber.Ctx) error {
	log.Println("GetAll handler")

	records, err := h.recordService.GetAllRecords(ctx.Context())
	if err != nil {
		log.Println("failed to get all records:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}
