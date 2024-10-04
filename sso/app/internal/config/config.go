package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

var config *Config

const defaultConfigPath = "config/config.yaml"

type Config struct {
	Env string `yaml:"env" env-default:"local"`
}

func GetConfig() *Config {
	return config
}

func MustLoad() *Config {
	return MustLoadPath(defaultConfigPath)
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	config = &Config{}
	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return config
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
