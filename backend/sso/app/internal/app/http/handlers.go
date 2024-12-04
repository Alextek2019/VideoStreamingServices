package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"vss/sso/internal/transport/http/auth"

	"vss/sso/internal/transport/http/user"

	"vss/sso/internal/config"
	"vss/sso/internal/service"
	logger "vss/sso/pkg/logger/handlers/slogpretty"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

type Services struct {
	userService         service.User
	authProviderService service.Provider
}

func (s *Server) MapHandlers(services Services) {
	s.mapMetrics()

	s.fiber.Use(cors.New(cors.Config{
		AllowOrigins:     "*", //     fmt.Sprintf("%s:%s", config.Get().Service.AllowedHost, config.Get().Service.AllowedPort),
		AllowHeaders:     "Accept,Accept-Language,Content-Language,Content-Type,fingerprint,User-Agent",
		AllowCredentials: false,
	}))

	s.userHandler = user.NewUserHandler(services.userService)
	s.authProviderHandler = auth.NewProviderHandler(services.authProviderService)

	groupApi := s.fiber.Group("api")
	groupV1 := groupApi.Group("v1")
	groupUser := groupV1.Group("user")
	groupAuth := groupV1.Group("auth")

	MapUserRoutes(groupUser, s.userHandler)
	MapAuthProviderRoutes(groupAuth, s.authProviderHandler)
}

func (s *Server) mapMetrics() {
	metricsApp := fiber.New(fiber.Config{DisableStartupMessage: true})
	prometheus := fiberprometheus.New(config.Get().Service.Name)
	prometheus.RegisterAt(metricsApp, "/metrics")

	s.fiber.Use(prometheus.Middleware)
	go func() {
		if err := metricsApp.Listen(fmt.Sprintf(":%s", config.Get().Metrics.Port)); err != nil {
			logger.Log.Error(err.Error())
		}
	}()
}
