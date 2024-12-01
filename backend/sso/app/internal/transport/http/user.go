package http

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	RegisterUser() fiber.Handler
	GetUser() fiber.Handler
	UpdateUser() fiber.Handler
	DeleteUser() fiber.Handler
}
