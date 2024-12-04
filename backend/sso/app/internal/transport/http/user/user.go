package user

import (
	domain "vss/sso/internal/domain/user"
	"vss/sso/internal/service"
	"vss/sso/internal/transport"
	"vss/sso/pkg/errors"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
	"vss/sso/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userService service.User
}

func NewUserHandler(userService service.User) transport.UserHandler {
	return &Handler{userService: userService}
}

func (u *Handler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params domain.RegisterUserArgs
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			logger.Log.Errorf("failed to validate body params: %v", err)
			return errors.ErrBodyParsing.ToFiberError(c)
		}

		usr, err := u.userService.Register(c.Context(), params)
		if err != nil {
			logger.Log.Errorf("failed to register user: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(usr)
	}
}

// TODO: implement
func (u *Handler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "unimplemented handler"})
	}
}
