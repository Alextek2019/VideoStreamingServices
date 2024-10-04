package app

import (
	"context"
	"github.com/pkg/errors"
	"vss/sso/internal/app/http"
)

type App struct {
	httpServer *http.Server
}

func New(ctx context.Context) (*App, error) {
	httpServer, err := http.NewServer(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "could not start http server")
	}

	return &App{
		httpServer: httpServer,
	}, nil
}

func (a *App) MustRun() {
	a.httpServer.MustRun()
}

func (a *App) Stop() {
	a.httpServer.Stop()
}
