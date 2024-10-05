package main

import (
	"context"
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

	application, err := app.New(context.Background())
	if err != nil {
		logger.Log.With("error", err.Error()).Error("could not start Application")
		return
	}

	go func() {
		application.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	logger.Log.Info("Gracefully stopped SSO")
}
