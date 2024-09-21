package helper

import "github.com/gofiber/fiber/v2"

type HelperInstance interface {
	Validate(interface{}) []ValidationError
	Response(*fiber.Ctx, int, bool, string, interface{}) error
	HandleValidationError([]ValidationError) string
}
