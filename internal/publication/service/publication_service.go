package service

import "github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/model"

type PublicationServiceInstance interface {
	IsPublisherExists(int) error

	IsPublicationExists(int) error
	IsPublicationsExist() error
	IsPublicationsExistByCategoryId(int) error
	IsPublicationCategoryExists(int) error

	GetPublications() ([]model.Publication, error)
	GetPublicationById(int) (model.Publication, error)
	GetPublicationsByCategoryId(int) ([]model.Publication, error)
	AddPublication(model.Publication) error
	UpdatePublication(model.UpdatedPublication) error
	DeletePublicationById(int) error

	IsPaperExists(int) error
	IsPapersExist() error
	IsOwnerExists(int) error
	GetPapers() ([]model.Paper, error)
	GetPaperById(int) (model.Paper, error)
	AddPaper(model.Paper) error
	UpdatePaper(model.UpdatedPaper) error
	DeletePaperById(int) error

	IsPaperFragmentExists(int) error
	AddPaperFragment(model.PaperFragment) error
	DeletePaperFragmentById(int) error

	IsCompetitionExists(int) error
	IsCompetitionCategoryExists(int) error
	IsCompetitionsExist() error
	IsCompetitionsExistByCategoryId(int) error
	GetCompetitions() ([]model.Competition, error)
	GetCompetitionById(int) (model.Competition, error)
	GetCompetitionsByCategoryId(int) ([]model.Competition, error)
	AddCompetition(model.Competition) error
	UpdateCompetition(model.UpdatedCompetition) error
	DeleteCompetitionById(int) error

	IsPublicationRequestExists(int) error
	DeletePublicationRequestById(int) error
}
