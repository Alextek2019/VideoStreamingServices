package http

import (
	"github.com/gofiber/fiber/v2"
	"vss/sso/internal/transport/http/middleware"
	"vss/sso/internal/transport/http/user"
)

func MapUserRoutes(userRoutes fiber.Router, h user.Handler, mw *middleware.MDWManager) {
	userRoutes.Post("/register", mw.UnAuthedMiddleware(), h.RegisterUser())
	userRoutes.Get("/get", mw.UnAuthedMiddleware(), h.RegisterUser())
	userRoutes.Patch("/update", mw.UnAuthedMiddleware(), h.RegisterUser())
	userRoutes.Delete("/delete", mw.UnAuthedMiddleware(), h.RegisterUser())
}
