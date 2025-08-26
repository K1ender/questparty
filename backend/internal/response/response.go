package response

import "github.com/gofiber/fiber/v3"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse(success bool, message string, data any) *Response {
	return &Response{Success: success, Message: message, Data: data}
}

func NewOkResponse(c fiber.Ctx, message string, data any) error {
	return c.Status(fiber.StatusOK).JSON(NewResponse(true, message, data))
}

func NewBadRequestResponse(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(NewResponse(false, message, nil))
}

func NewInternalServerErrorResponse(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(NewResponse(false, message, nil))
}
