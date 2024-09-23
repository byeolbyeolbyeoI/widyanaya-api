package handler

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/service"
)

type PublicationHandler struct {
	service service.UserServiceInstance
	helper  helper.HelperInstance
}

func NewPublicationHandler(s service.UserServiceInstance, h helper.HelperInstance) PublicationHandlerInstance {
	return &PublicationHandler{
		service: s,
		helper:  h,
	}
}
