package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

var config *Config

const defaultConfigPath = "config/config.yaml"

type Config struct {
	Env     string  `yaml:"Env" env-default:"local"`
	Service Service `yaml:"Service"`
	Metrics Metrics `yaml:"Metrics"`
}

type Service struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`

	AllowedHost string `yaml:"AllowedHost"`
	AllowedPort string `yaml:"AllowedPort"`
}

type Metrics struct {
	Port string `yaml:"Port"`
}

func Get() *Config {
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
