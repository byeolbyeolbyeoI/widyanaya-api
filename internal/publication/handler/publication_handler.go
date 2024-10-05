package handler

import "github.com/gofiber/fiber/v2"

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

	AddCompetition(c *fiber.Ctx) error
}
