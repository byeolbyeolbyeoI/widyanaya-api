package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty?
}

func (h *Helper) Response(c *fiber.Ctx, code int, s bool, msg string, data interface{}) error {
	return c.Status(code).JSON(&Response{
		Success: s,
		Message: msg,
		Data:    data,
	})
}
