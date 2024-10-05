package repository

import "github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/model"

type PublicationRepositoryInstance interface {
	IsPublisherExists(int) error

	IsPublicationsExist() error
	IsPublicationExists(int) error
	IsPublicationsExistByCategoryId(int) error
	IsPublicationCategoryExists(int) error

	GetPublications() ([]model.Publication, error)
	GetPublicationsByCategoryId(int) ([]model.Publication, error)
	GetPublicationById(int) (model.Publication, error)
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
	AddCompetition(model.Competition) error
	DeleteCompetitionById(int) error

	IsPublicationRequestExists(int) error
	DeletePublicationRequestById(int) error
}
