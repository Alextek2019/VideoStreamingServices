package middleware

import "github.com/gofiber/fiber/v2"

type MDWManager struct {
}

func (m *MDWManager) UnAuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}
