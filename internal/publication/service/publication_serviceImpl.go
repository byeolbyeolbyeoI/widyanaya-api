package service

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/repository"
)

type PublicationService struct {
	repo   repository.UserRepositoryInstance
	helper helper.HelperInstance
}

func NewPublicationService(r repository.UserRepositoryInstance, h helper.HelperInstance) PublicationServiceInstance {
	return &PublicationService{
		repo:   r,
		helper: h,
	}
}

func (u *PublicationService) CreatePublication() error {
	return nil
}
