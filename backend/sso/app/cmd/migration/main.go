package main

import (
	"flag"
	"vss/sso/internal/config"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
	"vss/sso/pkg/migrator"
)

const (
	actionFlagMigrationsUp            = "up"
	actionFlagMigrationsDown          = "down"
	actionFlagMigrationsUpOne         = "upone"
	acionFlagMigrationsGetCurrVersion = "version"
)

func main() {
	var action, cfgPath string

	flag.StringVar(&action, "action", "", "action to complete")
	flag.StringVar(&cfgPath, "cfg-path", "", "path to config")
	flag.Parse()

	if action == "" {
		panic("action arg is required")
	}

	cfg := config.MustLoad()

	logger.SetupLogger(cfg.Env)
	logger.Log.With("Env", cfg.Env).Info("Starting Migration For SSO")

	migr, err := migrator.New(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	switch action {
	case actionFlagMigrationsUp:
		if err := migr.Up(); err != nil {
			panic(err)
		}
	case actionFlagMigrationsDown:
		if err := migr.Down(); err != nil {
			panic(err)
		}
	case actionFlagMigrationsUpOne:
		if err := migr.UpByOne(); err != nil {
			panic(err)
		}
	case acionFlagMigrationsGetCurrVersion:
		if ver, err := migr.Version(); err != nil {
			panic(err)
		} else {
			logger.Log.Info("current db version %d", ver)
		}
	default:
		panic("unknown action flag: " + action)
	}
}
