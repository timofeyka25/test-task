package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"test-task/internal/dto"
	"test-task/internal/services"
	"test-task/pkg/validator"
)

type authHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *authHandler {
	return &authHandler{authService: authService}
}

func (h *authHandler) Register(router fiber.Router) {
	router.Post("/sign-up", h.signUp)
	router.Post("/sign-in", h.signIn)
}

func (h *authHandler) signUp(c *fiber.Ctx) error {
	log.Println("SignUp handler")

	var req dto.SignUpRequest

	if err := c.BodyParser(&req); err != nil {
		log.Println("failed to parse body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validator.ValidateStruct(&req); err != nil {
		log.Println("failed to validate:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error: Username should be longer than 6 characters, password longer than 8",
		})
	}

	resp, err := h.authService.SignUp(c.Context(), req)
	if err != nil {
		log.Println("failed to sign up:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to sign up user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *authHandler) signIn(c *fiber.Ctx) error {
	log.Println("SignIn handler")

	var req dto.SignInRequest

	if err := c.BodyParser(&req); err != nil {
		log.Println("failed to parse body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validator.ValidateStruct(&req); err != nil {
		log.Println("failed to validate:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error: Username should be longer than 6 characters, password longer than 8",
		})
	}

	resp, err := h.authService.SignIn(c.Context(), req)
	if err != nil {
		log.Println("failed to sign in:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
