package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var config Configuration

func Get() Configuration {
	return config
}

type Configuration struct {
}

func MustReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.AddConfigPath(params.Path.DirectoryPath)
	v.SetConfigName(params.Path.Name)
	v.SetConfigType(params.Path.Extension)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return err
	}

	if err = validator.New().Struct(tmp); err != nil {
		return err
	}

	c = &tmp

}
