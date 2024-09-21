package repository

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/model"
	supa "github.com/nedpals/supabase-go"
)

type UserRepository struct {
	client *supa.Client
	helper helper.HelperInstance
}

func NewUserRepository(s *supa.Client, h helper.HelperInstance) UserRepositoryInstance {
	return &UserRepository{
		client: s,
		helper: h,
	}
}

func (u *UserRepository) CreateUser(user model.User) error {
	var empty []map[string]interface{}

	err := u.client.DB.From("users").Insert(map[string]interface{}{
		"username":      user.Username,
		"password_hash": user.Password,
		"email":         user.Email,
	}).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}
