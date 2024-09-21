package service

import "github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/model"

type UserServiceInstance interface {
	HashPassword(string) (string, error)
	CreateUser(model.User) error
}
