package handler

import (
	"github.com/gofiber/fiber/v2"
)

type PublicationHandlerInstance interface {
	GetPublications(c *fiber.Ctx) error
	GetPublicationById(c *fiber.Ctx) error
	GetPublicationsByCategoryId(c *fiber.Ctx) error
	AddPublication(c *fiber.Ctx) error
	UpdatePublication(c *fiber.Ctx) error
	DeletePublicationById(c *fiber.Ctx) error

	GetPapers(c *fiber.Ctx) error
	GetPaperById(c *fiber.Ctx) error
	AddPaper(c *fiber.Ctx) error
	UpdatePaper(c *fiber.Ctx) error
	DeletePaperById(c *fiber.Ctx) error

	GetCompetitions(c *fiber.Ctx) error
	GetCompetitionById(c *fiber.Ctx) error
	GetCompetitionsByCategoryId(c *fiber.Ctx) error
	AddCompetition(c *fiber.Ctx) error
	UpdateCompetition(c *fiber.Ctx) error
	DeleteCompetitionById(c *fiber.Ctx) error

	GetPaperFragments(c *fiber.Ctx) error
	GetPaperFragmentById(c *fiber.Ctx) error
	AddPaperFragment(c *fiber.Ctx) error
	UpdatePaperFragment(c *fiber.Ctx) error
	DeletePaperFragmentById(c *fiber.Ctx) error

	GetPublicationRequests(c *fiber.Ctx) error
	GetPublicationRequestById(c *fiber.Ctx) error
	AddPublicationRequest(c *fiber.Ctx) error
	UpdatePublicationRequest(c *fiber.Ctx) error
	DeletePublicationRequestById(c *fiber.Ctx) error

	GetMetadatas(c *fiber.Ctx) error
	GetMetadataById(c *fiber.Ctx) error
	AddMetadata(c *fiber.Ctx) error
	UpdateMetadata(c *fiber.Ctx) error
	DeleteMetadataById(c *fiber.Ctx) error
}
