package auth

import (
	domain "vss/sso/internal/domain/auth"
	"vss/sso/internal/service"
	"vss/sso/internal/transport"
	logger "vss/sso/pkg/logger/handlers/slogpretty"

	"github.com/gofiber/fiber/v2"
)

type ProviderHandler struct {
	srv service.Provider
}

func NewProviderHandler(service service.Provider) transport.AuthProvider {
	return &ProviderHandler{
		srv: service,
	}
}

func (h *ProviderHandler) SignIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request domain.SignInRequest
		if err := c.BodyParser(&request); err != nil {
			logger.Log.Debugf("could not parse sign in body, error: %v", err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		resp, err := h.srv.SignIn(c.Context(), request)
		if err != nil {
			logger.Log.Debugf("could not sign in, error: %v", err)
			return err
		}

		authCookie := new(fiber.Cookie)
		authCookie.Name = "access_token"
		authCookie.Value = resp.AccessToken.String()
		authCookie.Expires = resp.TTL
		authCookie.HTTPOnly = true
		authCookie.Secure = true

		userCookie := new(fiber.Cookie)
		userCookie.Name = "user_id"
		userCookie.Value = resp.UserID.String()
		userCookie.HTTPOnly = true
		userCookie.Secure = true

		c.Cookie(authCookie)
		c.Cookie(userCookie)
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (h *ProviderHandler) SignOut() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.SendStatus(fiber.StatusOK)
	}
}
