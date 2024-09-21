package repository

import (
	"fmt"
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

func (u *UserRepository) IsExist(username string) (bool, error) {
	var result []struct {
		Username string `json:"username"`
	}
	err := u.client.DB.From("users").
		Select("username").
		Filter("username", "eq", username).
		Execute(&result)
	if err != nil {
		return false, err
	}

	if len(result) > 0 {
		return true, nil
	}

	return false, nil
}

func (u *UserRepository) GetPassword(username string) (string, error) {
	var result []struct {
		PasswordHash string `json:"password_hash"`
	}
	err := u.client.DB.From("users").
		Select("password_hash").
		Filter("username", "eq", username).
		Execute(&result)
	if err != nil {
		return "", err
	}

	// better error handling please
	// something like fmt.Errorf() n then the upper layer checks with error.Is() (kind of)
	// for now this aight
	if len(result) == 0 {
		fmt.Println(err.Error())
		return "", nil
	}

	return result[0].PasswordHash, nil
}
