package service

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/model"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/repository"
)

type PublicationService struct {
	repo   repository.PublicationRepositoryInstance
	helper helper.HelperInstance
}

func NewPublicationService(r repository.PublicationRepositoryInstance, h helper.HelperInstance) PublicationServiceInstance {
	return &PublicationService{
		repo:   r,
		helper: h,
	}
}

func (u *PublicationService) IsPublisherExists(id int) error {
	err := u.repo.IsPublisherExists(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *PublicationService) IsPublicationsExist() error {
	err := u.repo.IsPublicationsExist()
	if err != nil {
		return err
	}
	return nil
}

func (u *PublicationService) IsPublicationCategoryExists(categoryId int) error {
	err := u.repo.IsPublicationCategoryExists(categoryId)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsPublicationsExistByCategoryId(categoryId int) error {
	err := u.repo.IsPublicationsExistByCategoryId(categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (u *PublicationService) IsPublicationExists(id int) error {
	err := u.repo.IsPublicationExists(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) GetPublications() ([]model.Publication, error) {
	publications, err := u.repo.GetPublications()
	if err != nil {
		return nil, err
	}

	return publications, nil
}

func (u *PublicationService) GetPublicationById(id int) (model.Publication, error) {
	publication, err := u.repo.GetPublicationById(id)
	if err != nil {
		return model.Publication{}, err
	}

	return publication, nil
}

func (u *PublicationService) GetPublicationsByCategoryId(categoryId int) ([]model.Publication, error) {
	publications, err := u.repo.GetPublicationsByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}

	return publications, nil
}

func (u *PublicationService) AddPublication(publication model.Publication) error {
	err := u.repo.AddPublication(publication)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) UpdatePublication(publication model.UpdatedPublication) error {
	err := u.repo.UpdatePublication(publication)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) DeletePublicationById(id int) error {
	err := u.repo.DeletePublicationById(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsPaperExists(id int) error {
	err := u.repo.IsPaperExists(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsPapersExist() error {
	err := u.repo.IsPapersExist()
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) GetPapers() ([]model.Paper, error) {
	papers, err := u.repo.GetPapers()
	if err != nil {
		return nil, err
	}

	return papers, nil
}

func (u *PublicationService) GetPaperById(id int) (model.Paper, error) {
	paper, err := u.repo.GetPaperById(id)
	if err != nil {
		return model.Paper{}, err
	}

	return paper, nil
}

func (u *PublicationService) AddPaper(paper model.Paper) error {
	err := u.repo.AddPaper(paper)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsOwnerExists(id int) error {
	err := u.repo.IsOwnerExists(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *PublicationService) UpdatePaper(paper model.UpdatedPaper) error {
	err := u.repo.UpdatePaper(paper)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) DeletePaperById(id int) error {
	err := u.repo.DeletePaperById(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsPaperFragmentExists(id int) error {
	err := u.repo.IsPaperFragmentExists(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) AddPaperFragment(fragment model.PaperFragment) error {
	err := u.repo.AddPaperFragment(fragment)
	if err != nil {
		return err
	}

	return nil
}
func (u *PublicationService) DeletePaperFragmentById(id int) error {
	err := u.repo.DeletePaperFragmentById(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsCompetitionExists(id int) error {
	err := u.repo.IsCompetitionExists(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsCompetitionsExist() error {
	err := u.repo.IsCompetitionsExist()
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) GetCompetitions() ([]model.Competition, error) {
	competitions, err := u.repo.GetCompetitions()
	if err != nil {
		return nil, err
	}

	return competitions, nil
}
func (u *PublicationService) GetCompetitionById(id int) (model.Competition, error) {
	competition, err := u.repo.GetCompetitionById(id)
	if err != nil {
		return model.Competition{}, err
	}

	return competition, nil
}

func (u *PublicationService) IsCompetitionCategoryExists(categoryId int) error {
	err := u.repo.IsCompetitionCategoryExists(categoryId)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsCompetitionsExistByCategoryId(categoryId int) error {
	err := u.repo.IsCompetitionsExistByCategoryId(categoryId)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) GetCompetitionsByCategoryId(categoryId int) ([]model.Competition, error) {
	competitions, err := u.repo.GetCompetitionsByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}

	return competitions, nil
}

func (u *PublicationService) AddCompetition(competition model.Competition) error {
	err := u.repo.AddCompetition(competition)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) UpdateCompetition(competition model.UpdatedCompetition) error {
	err := u.repo.UpdateCompetition(competition)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) DeleteCompetitionById(id int) error {
	err := u.repo.DeleteCompetitionById(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) IsPublicationRequestExists(id int) error {
	err := u.repo.IsPublicationRequestExists(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationService) DeletePublicationRequestById(id int) error {
	err := u.repo.DeletePublicationRequestById(id)
	if err != nil {
		return err
	}

	return nil
}
