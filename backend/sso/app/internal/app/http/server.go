package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"vss/sso/internal/config"
	"vss/sso/internal/service/auth"
	userService "vss/sso/internal/service/user"
	"vss/sso/internal/transport"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
)

type Server struct {
	fiber *fiber.App

	userHandler         transport.UserHandler
	authProviderHandler transport.AuthProvider

	ctx context.Context
}

func NewServer(ctx context.Context) (*Server, error) {
	srv := Server{
		fiber: fiber.New(fiber.Config{DisableStartupMessage: true, EnableSplittingOnParsers: true}),
	}

	userSvc, err := userService.New(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "app.http.NewServer could not create user service")
	}

	authProviderSvc, err := auth.NewProvider(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "app.http.NewServer could not create auth service")
	}

	srv.MapHandlers(Services{
		userService:         userSvc,
		authProviderService: authProviderSvc,
	})

	return &srv, nil
}

func (s *Server) MustRun() {
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", config.Get().Service.Host, config.Get().Service.Port)); err != nil {
		panic(err)
	}
}

func (s *Server) Stop() {
	err := s.fiber.Shutdown()
	if err != nil {
		logger.Log.Errorf("couldn`t gracefully stop fiber, error: %v", err)
	}
}
