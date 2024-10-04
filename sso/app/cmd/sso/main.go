package main

import (
	"os"
	"os/signal"
	"syscall"
	"vss/sso/internal/app"
	logger "vss/sso/pkg/logger/handlers/slogpretty"

	"vss/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	logger.SetupLogger(cfg.Env)
	logger.Log.Info("Starting SSO Service")

	application := app.New()
	go func() {
		application.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.MustStop()
	logger.Log.Info("Gracefully stopped SSO")
}
