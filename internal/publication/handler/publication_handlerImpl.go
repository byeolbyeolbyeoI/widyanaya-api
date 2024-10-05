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

func (p *PublicationHandler) GetCompetitions(c *fiber.Ctx) error {
	err := p.service.IsCompetitionsExist()
	if err != nil {
		if errors.Is(err, helper.ErrCompetitionNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	competitions, err := p.service.GetCompetitions()
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	sort.Slice(competitions, func(i int, j int) bool {
		return competitions[i].OpeningDate.Before(competitions[j].OpeningDate)
	})

	return p.helper.Response(c, fiber.StatusOK, true, "competitions retrieved successfully", competitions)
}

func (p *PublicationHandler) GetCompetitionById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsCompetitionExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrCompetitionNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	competition, err := p.service.GetCompetitionById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "competition retrieved successfully", competition)
}

func (p *PublicationHandler) GetCompetitionsByCategoryId(c *fiber.Ctx) error {
	strCategoryId := c.Params("category_id")
	categoryId, err := strconv.Atoi(strCategoryId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsCompetitionCategoryExists(categoryId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
	}

	err = p.service.IsCompetitionsExistByCategoryId(categoryId)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	competitions, err := p.service.GetCompetitionsByCategoryId(categoryId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	sort.Slice(competitions, func(i int, j int) bool {
		return competitions[i].OpeningDate.Before(competitions[j].OpeningDate)
	})

	return p.helper.Response(c, fiber.StatusOK, true, "competitions retrieved successfully", competitions)
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

	err = p.service.IsCompetitionCategoryExists(competition.CompetitionCategoryId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
	}

	err = p.service.IsPublisherExists(competition.PublisherId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
	}

	err = p.service.AddCompetition(competition)
	if err != nil {
		return p.helper.Response(c, fiber.StatusBadRequest, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "competition added successfully", nil)
}

func (p *PublicationHandler) UpdateCompetition(c *fiber.Ctx) error {
	var competition model.UpdatedCompetition
	err := c.BodyParser(&competition)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(competition); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsCompetitionExists(competition.Id)
	if err != nil {
		if errors.Is(err, helper.ErrCompetitionNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.UpdateCompetition(competition)
	if err != nil {
		return p.helper.Response(c, fiber.StatusBadRequest, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "competition updated successfully", nil)
}

func (p *PublicationHandler) DeleteCompetitionById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	err = p.service.IsCompetitionExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrCompetitionNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.DeleteCompetitionById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "competition deleted successfully", nil)
}

func (p *PublicationHandler) GetPaperFragments(c *fiber.Ctx) error {
	err := p.service.IsPaperFragmentsExist()
	if err != nil {
		if errors.Is(err, helper.ErrPaperFragmentNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	fragments, err := p.service.GetPaperFragments()
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper fragments retrieved successfully", fragments)
}

func (p *PublicationHandler) GetPaperFragmentById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}
	err = p.service.IsPaperFragmentExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPaperFragmentNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	fragment, err := p.service.GetPaperFragmentById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "fragment retrieved successfully", fragment)
}

func (p *PublicationHandler) AddPaperFragment(c *fiber.Ctx) error {
	var fragment model.PaperFragment
	err := c.BodyParser(&fragment)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(fragment); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsPaperExists(fragment.PaperId)
	if err != nil {
		if errors.Is(err, helper.ErrPaperNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return err
	}

	err = p.service.IsPaperFragmentCategoryExists(fragment.FragmentCategoryId)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return err
	}

	// cek if exist if ya if if
	err = p.service.AddPaperFragment(fragment)
	if err != nil { // better err handler please
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper fragment added successfully", nil)
}

func (p *PublicationHandler) UpdatePaperFragment(c *fiber.Ctx) error {
	var fragment model.UpdatedPaperFragment
	err := c.BodyParser(&fragment)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(fragment); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsPaperFragmentExists(fragment.Id)
	if err != nil {
		if errors.Is(err, helper.ErrPaperFragmentNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsPaperExists(fragment.PaperId)
	if err != nil {
		if errors.Is(err, helper.ErrPaperNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return err
	}

	err = p.service.IsPaperFragmentCategoryExists(fragment.FragmentCategoryId)
	if err != nil {
		if errors.Is(err, helper.ErrCategoryNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	// cek if exist if ya if if
	err = p.service.UpdatePaperFragment(fragment)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper fragment updated successfully", fragment)
}

func (p *PublicationHandler) DeletePaperFragmentById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	err = p.service.IsPaperFragmentExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPaperFragmentNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.DeletePaperFragmentById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "paper fragment deleted successfully", nil)
}

func (p *PublicationHandler) GetPublicationRequests(c *fiber.Ctx) error {
	err := p.service.IsPublicationRequestsExist()
	if err != nil {
		if errors.Is(err, helper.ErrPublicationRequestNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	requests, err := p.service.GetPublicationRequests()
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication requests retrieved successfully", requests)
}

func (p *PublicationHandler) GetPublicationRequestById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsPublicationRequestExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationRequestNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	request, err := p.service.GetPublicationRequestById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication request retrieved successfully", request)
}

func (p *PublicationHandler) DeletePublicationRequestById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	err = p.service.IsPublicationRequestExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationRequestNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.DeletePublicationRequestById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication request deleted successfully", nil)
}

func (p *PublicationHandler) AddPublicationRequest(c *fiber.Ctx) error {
	var request model.PublicationRequest
	err := c.BodyParser(&request)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(request); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	// is metadata id
	// reference_format_id
	//requester id
	err = p.service.IsMetadataExists(request.MetadataID)
	if err != nil {
		if errors.Is(err, helper.ErrMetadataNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsReferenceFormatExists(request.ReferenceFormatID)
	if err != nil {
		if errors.Is(err, helper.ErrReferenceFormatNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsRequesterExists(request.RequesterID)
	if err != nil {
		if errors.Is(err, helper.ErrRequesterNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.AddPublicationRequest(request)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication request added successfully", request)

}

func (p *PublicationHandler) UpdatePublicationRequest(c *fiber.Ctx) error {
	var req model.UpdatedPublicationRequest
	err := c.BodyParser(&req)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(req); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsPublicationRequestExists(req.Id)
	if err != nil {
		if errors.Is(err, helper.ErrPublicationRequestNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsReferenceFormatExists(req.ReferenceFormatID)
	if err != nil {
		if errors.Is(err, helper.ErrReferenceFormatNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsMetadataExists(req.MetadataID)
	if err != nil {
		if errors.Is(err, helper.ErrMetadataNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsRequesterExists(req.RequesterID)
	if err != nil {
		if errors.Is(err, helper.ErrRequesterNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.UpdatePublicationRequest(req)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "publication request updated successfully", req)
}

func (p *PublicationHandler) DeleteMetadataById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	err = p.service.IsMetadataExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrMetadataNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.DeleteMetadataById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "metadata deleted successfully", nil)
}

func (p *PublicationHandler) GetMetadatas(c *fiber.Ctx) error {
	err := p.service.IsMetadatasExist()
	if err != nil {
		if errors.Is(err, helper.ErrMetadataNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	metadatas, err := p.service.GetMetadatas()
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "metadata retrieved successfully", metadatas)
}

func (p *PublicationHandler) GetMetadataById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.IsMetadataExists(id)
	if err != nil {
		if errors.Is(err, helper.ErrMetadataNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}

		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	metadata, err := p.service.GetMetadataById(id)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "metadata retrieved successfully", metadata)
}

func (p *PublicationHandler) AddMetadata(c *fiber.Ctx) error {
	var metadata model.Metadata
	err := c.BodyParser(&metadata)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(metadata); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	// cek if exist if ya if if
	err = p.service.AddMetadata(metadata)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "metadata added successfully", metadata)
}

func (p *PublicationHandler) UpdateMetadata(c *fiber.Ctx) error {
	var metadata model.UpdatedMetadata
	err := c.BodyParser(&metadata)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := p.helper.Validate(metadata); len(errs) > 0 && errs[0].Error {
		errMsg := p.helper.HandleValidationError(errs)

		return p.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	err = p.service.IsMetadataExists(metadata.Id)
	if err != nil {
		if errors.Is(err, helper.ErrMetadataNotFound) {
			return p.helper.Response(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = p.service.UpdateMetadata(metadata)
	if err != nil {
		return p.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return p.helper.Response(c, fiber.StatusOK, true, "metadata updated successfully", metadata)
}
