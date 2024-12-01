package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"vss/sso/internal/config"
	userService "vss/sso/internal/service/user"
	"vss/sso/internal/transport/http"
)

type Server struct {
	fiber *fiber.App

	userHandler http.UserHandler
	ctx         context.Context
}

func NewServer(ctx context.Context) (*Server, error) {
	srv := Server{
		fiber: fiber.New(fiber.Config{DisableStartupMessage: true, EnableSplittingOnParsers: true}),
	}

	userSvc, err := userService.New(ctx)
	if err != nil {
		return nil, errors.Wrapf(err,
			"app.http.NewServer %s",
			"could not create user service")
	}

	handlers := Services{
		UserService: userSvc,
	}

	err = srv.MapHandlers(handlers)
	if err != nil {
		return nil, errors.Wrapf(err,
			"app.http.NewServer %s",
			"could not map handlers")
	}

	return &srv, nil
}

func (s *Server) MustRun() {
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", config.Get().Service.Host, config.Get().Service.Port)); err != nil {
		panic(err)
	}
}

func (s *Server) Stop() {
	_ = s.fiber.Shutdown()
}
