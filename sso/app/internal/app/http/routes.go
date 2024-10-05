package http

import (
	"github.com/gofiber/fiber/v2"
	"vss/sso/internal/transport/http"
	"vss/sso/internal/transport/http/middleware"
)

func MapUserRoutes(userRoutes fiber.Router, h http.UserHandler, mw *middleware.MDWManager) {
	userRoutes.Post("/register", mw.UnAuthedMiddleware(), h.RegisterUser())
	userRoutes.Get("/get", mw.UnAuthedMiddleware(), h.RegisterUser())
	userRoutes.Patch("/update", mw.UnAuthedMiddleware(), h.RegisterUser())
	userRoutes.Delete("/delete", mw.UnAuthedMiddleware(), h.RegisterUser())
}
