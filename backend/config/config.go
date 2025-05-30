package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ApiServerPort string `env:"APISERVER_PORT"`
	ApiServerHost string `env:"APISERVER_HOST"`

	DatabaseName     string `env:"DB_NAME"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     string `env:"DB_PORT"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`

	DatabaseMaxConn           int32         `env:"DB_MAX_CONN"`
	DatabaseMinConn           int32         `env:"DB_MIN_CONN"`
	DatabaseMaxConnLifetime   time.Duration `env:"DB_MAX_CONN_LIFETIME"`
	DatabaseMaxConnIdletime   time.Duration `env:"DB_MAX_CONN_IDLETIME"`
	DatabaseHealthCheckPeriod time.Duration `env:"DB_HEALTH_CHECK_PERIOD"`

	JwtSecret string `env:"JWT_SECRET"`
}

func (c *Config) DatabaseUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
	)
}

func New() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}
	return &cfg, nil
}
