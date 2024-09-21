package service

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/model"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo   repository.UserRepositoryInstance
	helper helper.HelperInstance
}

func NewUserService(r repository.UserRepositoryInstance, h helper.HelperInstance) UserServiceInstance {
	return &UserService{
		repo:   r,
		helper: h,
	}
}

func (u *UserService) CreateUser(user model.User) error {
	err := u.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
