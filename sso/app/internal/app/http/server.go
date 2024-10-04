package http

import "vss/sso/internal/transport/http"

type App struct {
	userHandler http.UserHandler
}

func NewApp() *App {
	return &App{}
}
