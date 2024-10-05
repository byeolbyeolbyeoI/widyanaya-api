package handler

import (
	"errors"
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/model"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/service"
	"github.com/gofiber/fiber/v2"
	"sort"
	"strconv"
)

type PublicationHandler struct {
	service service.PublicationServiceInstance
	helper  helper.HelperInstance
}

func NewPublicationHandler(s service.PublicationServiceInstance, h helper.HelperInstance) PublicationHandlerInstance {
	return &PublicationHandler{
		service: s,
		helper:  h,
	}
}

func (p *PublicationHandler) GetPublications(c *fiber.Ctx) error {
	err := p.service.IsPublicationsExist()
	if err != nil {
		if errors.Is(err, helper.ErrPublicationNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	publications, err := p.service.GetPublications()
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	sort.Slice(publications, func(i int, j int) bool {
		return publications[i].OpeningDate.Before(publications[j].OpeningDate)
	})

	return p.helper.Response(c, fiber.StatusOK, true, "publications retrieved successfully", publications)
}

func (p *PublicationHandler) GetPublicationById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsPublicationExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	publication, err := p.service.GetPublicationById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication retrieved successfully", publication)
}

func (p *PublicationHandler) AddPublication(c *fiber.Ctx) error {
	var publication model.Publication
	err := c.BodyParser(&publication)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(publication); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsPublisherExists(publication.PublisherId)
	if err != nil {
		if errors.Is(err, helper.ErrPublisherNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return err
	}

	err = p.service.IsPublicationCategoryExists(publication.PublicationCategoryID)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}
	// cek if exist if ya if if
	err = p.service.AddPublication(publication)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication added successfully", publication)
}

func (p *PublicationHandler) GetPublicationsByCategoryId(c *fiber.Ctx) error {
	strCategoryId := c.Params("category_id")
	categoryId, err := strconv.Atoi(strCategoryId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsPublicationCategoryExists(categoryId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
	}

	err = p.service.IsPublicationsExistByCategoryId(categoryId)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	publications, err := p.service.GetPublicationsByCategoryId(categoryId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	sort.Slice(publications, func(i int, j int) bool {
		return publications[i].OpeningDate.Before(publications[j].OpeningDate)
	})

	return p.helper.Response(c, fiber.StatusOK, true, "publication retrieved successfully", publications)
}

func (p *PublicationHandler) UpdatePublication(c *fiber.Ctx) error {
	var publication model.UpdatedPublication
	err := c.BodyParser(&publication)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(publication); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsPublicationExists(publication.Id)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsPublisherExists(publication.PublisherId)
	if err != nil {
		if errors.Is(err, helper.ErrPublisherNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return err
	}

	err = p.service.IsPublicationCategoryExists(publication.PublicationCategoryID)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	// cek if exist if ya if if
	err = p.service.UpdatePublication(publication)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication updated successfully", publication)
}

func (p *PublicationHandler) DeletePublicationById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	err = p.service.IsPublicationExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.DeletePublicationById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication deleted successfully", nil)
}

func (p *PublicationHandler) GetPapers(c *fiber.Ctx) error {
	err := p.service.IsPapersExist()
	if err != nil {
		if errors.Is(err, helper.ErrPaperNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	papers, err := p.service.GetPapers()
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "papers retrieved successfully", papers)
}

func (p *PublicationHandler) GetPaperById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}
	err = p.service.IsPaperExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPaperNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	paper, err := p.service.GetPaperById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper retrieved successfully", paper)
}

func (p *PublicationHandler) AddPaper(c *fiber.Ctx) error {
	var paper model.Paper
	err := c.BodyParser(&paper)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(paper); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsOwnerExists(paper.OwnerId)
	if err != nil {
		if errors.Is(err, helper.ErrOwnerNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return err
	}
	// cek if exist if ya if if
	err = p.service.AddPaper(paper)
	if err != nil { // better err handler please

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper added successfully", nil)
}

func (p *PublicationHandler) UpdatePaper(c *fiber.Ctx) error {
	var paper model.UpdatedPaper
	err := c.BodyParser(&paper)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(paper); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsPaperExists(paper.Id)
	if err != nil {
		if errors.Is(err, helper.ErrPaperNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsOwnerExists(paper.OwnerId)
	if err != nil {
		if errors.Is(err, helper.ErrOwnerNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return err
	}

	// cek if exist if ya if if
	err = p.service.UpdatePaper(paper)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper updated successfully", paper)
}

func (p *PublicationHandler) DeletePaperById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	err = p.service.IsPaperExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPaperNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.DeletePaperById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper deleted successfully", nil)
}

func (p *PublicationHandler) AddCompetition(c *fiber.Ctx) error {
	var competition model.Competition
	err := c.BodyParser(&competition)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(competition); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.AddCompetition(competition)
	if err != nil {
		return p.helper.Response(c, fiber.StatusBadRequest, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "competition added successfully", nil)
}
