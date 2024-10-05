package pgconnector

import (
	"fmt"
	"time"
)

type Postgres struct {
	Host     string `yaml:"Host" validate:"required"`
	Port     string `yaml:"Port" validate:"required"`
	User     string `yaml:"User" validate:"required"`
	Password string `yaml:"Password" validate:"required"`
	DBName   string `yaml:"DBName" validate:"required"`
	SSLMode  string `yaml:"SSLMode" validate:"required"`
	PGDriver string `yaml:"PGDriver" validate:"required"`

	MaxOpenConns    int           `yaml:"MaxOpenConns" validate:"required,min=1"`
	ConnMaxLifetime time.Duration `yaml:"ConnMaxLifetime" validate:"required,min=1"`
	MaxIdleConns    int           `yaml:"MaxIdleConns" validate:"required,min=1"`
	ConnMaxIdleTime time.Duration `yaml:"ConnMaxIdleTime" validate:"required,min=1"`
}

func (c *Postgres) BuildConnectionUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
	)
}
