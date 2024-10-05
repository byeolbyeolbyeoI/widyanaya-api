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

	app.Get("/paper_fragments", publicationHandler.GetPaperFragments)
	app.Get("/paper_fragment/:id", publicationHandler.GetPaperFragmentById)
	app.Put("/paper_fragment", publicationHandler.UpdatePaperFragment)
	app.Post("/paper_fragment", publicationHandler.AddPaperFragment)
	app.Delete("/paper_fragment/:id", publicationHandler.DeletePaperFragmentById)

	app.Get("/competitions", publicationHandler.GetCompetitions)
	app.Get("/competitions/:category_id", publicationHandler.GetCompetitionsByCategoryId)
	app.Get("/competition/:id", publicationHandler.GetCompetitionById)
	app.Post("/competition", publicationHandler.AddCompetition)
	app.Put("/competition", publicationHandler.UpdateCompetition)
	app.Delete("/competition/:id", publicationHandler.DeleteCompetitionById)

	app.Get("/publication_requests", publicationHandler.GetPublicationRequests)
	app.Get("/publication_request/:id", publicationHandler.GetPublicationRequestById)
	app.Post("/publication_request", publicationHandler.AddPublicationRequest)
	app.Put("/publication_request", publicationHandler.UpdatePublicationRequest)
	app.Delete("/publication_request/:id", publicationHandler.DeletePublicationRequestById)

	app.Get("/metadatas", publicationHandler.GetMetadatas)
	app.Get("/metadata/:id", publicationHandler.GetMetadataById)
	app.Post("/metadata", publicationHandler.AddMetadata)
	app.Put("/metadata", publicationHandler.UpdateMetadata)
	app.Delete("/metadata/:id", publicationHandler.DeleteMetadataById)

	return app
}
