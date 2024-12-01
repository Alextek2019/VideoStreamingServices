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
	logger.Log.Infof("Server name %s", cfg.Service.Name)

	application, err := app.New(context.Background())
	if err != nil {
		logger.Log.Errorf("could not start Application, err: %s", err)
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
