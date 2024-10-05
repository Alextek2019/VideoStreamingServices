package user

import (
	"vss/sso/internal/service/user"
	"vss/sso/internal/transport/http"
	"vss/sso/pkg/errors"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
	"vss/sso/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userService *user.Service
}

func NewUserHandler(userService *user.Service) http.UserHandler {
	return &Handler{userService: userService}
}

func (u *Handler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params RegisterUserRequest
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			logger.Log.Error(
				"Error: %v customError: %v\nUser.Transport.http.RegisterUser()",
				err,
				errors.ErrBodyParsing.GetErrConst(),
			)
			return errors.ErrBodyParsing.ToFiberError(c)
		}

		return c.Status(fiber.StatusOK).JSON(nil)
	}
}

func (u *Handler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(nil)
	}
}

func (u *Handler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params RegisterUserRequest
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			logger.Log.Error(
				"Error: %v customError: %v\nUser.Transport.http.UpdateUser()",
				err,
				errors.ErrBodyParsing.GetErrConst(),
			)
			return errors.ErrBodyParsing.ToFiberError(c)
		}

		return c.Status(fiber.StatusOK).JSON(nil)
	}
}

func (u *Handler) DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(nil)
	}
}
