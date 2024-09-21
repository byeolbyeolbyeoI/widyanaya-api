package repository

import "github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/model"

type UserRepositoryInstance interface {
	CreateUser(model.User) error
	IsExist(string) (bool, error)
	GetPassword(string) (string, error)
}
