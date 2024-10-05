package handler

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/model"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/service"
	"github.com/gofiber/fiber/v2"
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
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := u.helper.Validate(user); len(errs) > 0 && errs[0].Error {
		errMsg := u.helper.HandleValidationError(errs)

		return u.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	exists, err := u.service.IsExist(user.Username)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if exists {
		return u.helper.Response(c, fiber.StatusConflict, false, "invalid username or password", nil)
	}

	user.PasswordHash, err = u.service.HashPassword(user.PasswordHash)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	err = u.service.CreateUser(user)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return u.helper.Response(c, fiber.StatusOK, true, "user signed up successfully", nil)
}

func (u *UserHandler) Login(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if errs := u.helper.Validate(user); len(errs) > 0 && errs[0].Error {
		errMsg := u.helper.HandleValidationError(errs)

		return u.helper.Response(c, fiber.StatusBadRequest, false, errMsg, nil)
	}

	exists, err := u.service.IsExist(user.Username)
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if !exists {
		return u.helper.Response(c, fiber.StatusConflict, false, "invalid username or password", nil)
	}

	err = u.service.CheckPassword(user.Username, user.PasswordHash)
	if err != nil { // pass salah
		return u.helper.Response(c, fiber.StatusUnauthorized, false, "invalid username or password", nil)
	}

	return u.helper.Response(c, fiber.StatusOK, true, "user logged in successfully", nil)
}
