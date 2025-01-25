package shared

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(&ApiResponse{
		Success: true,
		Message: message,
		Data: data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(&ApiResponse{
		Success: false,
		Message: message,
	})
}