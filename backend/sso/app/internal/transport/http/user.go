package http

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	RegisterUser() fiber.Handler
	UpdateUser() fiber.Handler
}
