package server

import (
	publicationHandler "github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/handler"
	userHandler "github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/handler"
	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(app *fiber.App, userHandler userHandler.UserHandlerInstance, publicationHandler publicationHandler.PublicationHandlerInstance) *fiber.App {
	app.Post("/signup", userHandler.SignUp)
	app.Post("/login", userHandler.Login)

	app.Get("/publications", publicationHandler.GetPublications)
	app.Get("/publications/:category_id", publicationHandler.GetPublicationsByCategoryId)
	app.Get("/publication/:id", publicationHandler.GetPublicationById)
	app.Put("/publication", publicationHandler.UpdatePublication)
	app.Post("/publication", publicationHandler.AddPublication)
	app.Delete("/publication/:id", publicationHandler.DeletePublicationById)

	app.Get("/papers", publicationHandler.GetPapers)
	app.Get("/paper/:id", publicationHandler.GetPaperById)
	app.Put("/paper", publicationHandler.UpdatePaper)
	app.Post("/paper", publicationHandler.AddPaper)
	app.Delete("/paper/:id", publicationHandler.DeletePaperById)

	app.Post("/competition", publicationHandler.AddCompetition)

	return app
}
