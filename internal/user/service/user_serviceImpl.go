package service

import (
	"fmt"
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

func (u *UserService) IsExist(username string) (bool, error) {
	exists, err := u.repo.IsExist(username)
	if err != nil {
		return false, err
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func (u *UserService) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (u *UserService) CheckPassword(username string, password string) error {
	hashedPassword, err := u.repo.GetPassword(username)
	if err != nil {
		return err
	}
	fmt.Println("password			:", password)
	fmt.Println("hashed password	:", hashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			fmt.Println("wrong password")
		}
		return err
	}

	return nil
}
