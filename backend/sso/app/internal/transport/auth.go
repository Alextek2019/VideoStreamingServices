package transport

import "github.com/gofiber/fiber/v2"

type AuthProvider interface {
	SignIn() fiber.Handler
	SignOut() fiber.Handler
}

type AuthValidator interface {
	VerifyAuthToken()
}
