package http

import (
	"github.com/gofiber/fiber/v2"
	"vss/sso/internal/transport"
)

func MapUserRoutes(userRoutes fiber.Router, h transport.UserHandler) {
	userRoutes.Post("/", h.RegisterUser())
	userRoutes.Patch("/", h.UpdateUser())
}

func MapAuthProviderRoutes(authRoutes fiber.Router, h transport.AuthProvider) {
	authRoutes.Post("/signin", h.SignIn())
	authRoutes.Post("/signout", h.SignOut())
}
