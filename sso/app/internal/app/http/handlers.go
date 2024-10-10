package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"vss/sso/internal/transport/http/middleware"
	"vss/sso/internal/transport/http/user"

	"vss/sso/internal/config"
	"vss/sso/internal/service"
	logger "vss/sso/pkg/logger/handlers/slogpretty"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

type Services struct {
	UserService service.User
}

func (s *Server) MapHandlers(services Services) error {
	s.mapMetrics()

	s.fiber.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5000",
		AllowHeaders:     "Accept,Accept-Language,Content-Language,Content-Type,fingerprint,User-Agent",
		AllowCredentials: true,
	}))

	s.userHandler = user.NewUserHandler(services.UserService)

	mw := middleware.NewMDWManager()

	apiGroup := s.fiber.Group("api")
	v1 := apiGroup.Group("v1")
	userGroup := v1.Group("user")

	MapUserRoutes(userGroup, s.userHandler, mw)

	return nil
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
