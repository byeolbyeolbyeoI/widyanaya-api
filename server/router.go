package server

import (
	userHandler "github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/handler"
	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(app *fiber.App, userHandler userHandler.UserHandlerInstance) *fiber.App {
	app.Post("/signup", userHandler.SignUp)

	return app
}
