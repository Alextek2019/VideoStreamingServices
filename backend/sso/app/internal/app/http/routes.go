package http

import (
	"github.com/gofiber/fiber/v2"
	"vss/sso/internal/transport/http"
	"vss/sso/internal/transport/http/middleware"
)

func MapUserRoutes(userRoutes fiber.Router, h http.UserHandler, mw *middleware.MDWManager) {
	userRoutes.Post("/", mw.UnAuthedMiddleware(), h.RegisterUser())
	userRoutes.Get("/", mw.UnAuthedMiddleware(), h.GetUser())
	userRoutes.Patch("/", mw.UnAuthedMiddleware(), h.UpdateUser())
	userRoutes.Delete("/", mw.UnAuthedMiddleware(), h.DeleteUser())
}