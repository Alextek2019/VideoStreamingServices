package user

import (
	"github.com/gofrs/uuid"
	domain "vss/sso/internal/domain/user"
	"vss/sso/internal/service"
	"vss/sso/internal/transport/http"
	"vss/sso/pkg/errors"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
	"vss/sso/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userService service.User
}

func NewUserHandler(userService service.User) http.UserHandler {
	return &Handler{userService: userService}
}

func (u *Handler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params domain.RegisterUserArgs
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			logger.Log.With("error:", errors.ErrBodyParsing).Error("User.Transport.http.RegisterUserArgs")
			return errors.ErrBodyParsing.ToFiberError(c)
		}

		err := u.userService.Register(c.Context(), params)
		if err != nil {
			logger.Log.Error(
				"User.Transport.http.RegisterUserArgs",
				err,
				errors.ErrBodyParsing.GetErrConst(),
			)
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(nil)
	}
}

func (u *Handler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Query("userID")
		userUUID, err := uuid.FromString(userID)
		if err != nil {
			logger.Log.Error(
				"User.Transport.http.GetUser",
				err,
				errors.ErrUserID.GetErrConst(),
			)
			return errors.ErrUserID.ToFiberError(c)
		}

		userEntity, err := u.userService.Get(c.Context(), userUUID)
		if err != nil {
			logger.Log.Error(
				"User.Transport.http.GetUser",
				err,
				errors.ErrUserNotFound.GetErrConst(),
			)
			return errors.ErrUserNotFound.ToFiberError(c)
		}

		return c.Status(fiber.StatusOK).JSON(userEntity)
	}
}

func (u *Handler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(nil)
	}
}

func (u *Handler) DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(nil)
	}
}
