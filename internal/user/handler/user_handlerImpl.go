package handler

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/config"
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/model"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserHandler struct {
	service service.UserServiceInstance
	helper  helper.HelperInstance
	config  *config.Config
}

func NewUserHandler(s service.UserServiceInstance, h helper.HelperInstance, c *config.Config) UserHandlerInstance {
	return &UserHandler{
		service: s,
		helper:  h,
		config:  c,
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
	var user model.UserCredential
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud":     "widyanaya-api",
		"iss":     "widyanaya",
		"subject": user.Username,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.config.JWT.Secret))
	if err != nil {
		return u.helper.Response(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return u.helper.Response(c, fiber.StatusOK, true, "user logged in successfully", tokenString)
}
