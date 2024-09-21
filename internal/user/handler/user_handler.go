package handler

import "github.com/gofiber/fiber/v2"

type UserHandlerInstance interface {
	SignUp(*fiber.Ctx) error
}
