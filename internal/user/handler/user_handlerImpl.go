package handler

import (
	"fmt"
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/model"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type UserHandler struct {
	service service.UserServiceInstance
	helper  helper.HelperInstance
}

func NewUserHandler(s service.UserServiceInstance, h helper.HelperInstance) UserHandlerInstance {
	return &UserHandler{
		service: s,
		helper:  h,
	}
}

func (u *UserHandler) SignUp(c *fiber.Ctx) error {
	// take input
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	// validate it
	if errs := u.helper.Validate(user); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, e := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | needs to implement '%s'",
				e.FailedField,
				e.Value,
				e.Tag,
			))
		}

		return u.helper.Response(c, fiber.StatusBadRequest, false, strings.Join(errMsgs, " and "), nil)
	}

	// check if exists, later

	// hash it
	user.Password, err = u.service.HashPassword(user.Password)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	// create it
	err = u.service.CreateUser(user)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return u.helper.Response(c, fiber.StatusOK, true, "user signed up successfully", nil)
}
