package repository

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	supa "github.com/nedpals/supabase-go"
)

type PublicationRepository struct {
	client *supa.Client
	helper helper.HelperInstance
}

func NewPublicationRepository(s *supa.Client, h helper.HelperInstance) PublicationRepositoryInstance {
	return &PublicationRepository{
		client: s,
		helper: h,
	}
}

// refactor struct declare tiap handler

func (u *PublicationRepository) CreatePublication() error {
	return nil
}
