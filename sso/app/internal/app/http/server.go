package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"vss/sso/internal/config"
	userService "vss/sso/internal/service/user"
	"vss/sso/internal/transport/http"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
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

	handlers := Services{
		UserService: userService.New(ctx),
	}

	err := srv.MapHandlers(ctx, handlers)
	if err != nil {
		return nil, errors.Wrapf(err, "app.http.NewServer could not map handlers")
	}

	return &Server{}, nil
}

func (s *Server) MustRun() {
	logger.Log.With("Service", config.Get().Service).Info("Start http server")
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", config.Get().Service.Host, config.Get().Service.Port)); err != nil {
		logger.Log.With("error", err.Error()).Error("Error starting http Server")
		panic(err)
	}
}

func (s *Server) Stop() {
	logger.Log.Info("stopping http server")
	_ = s.fiber.Shutdown()
}
