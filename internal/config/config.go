package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig     DBConfig
	ServerConfig ServerConfig
	LoggerConfig LoggerConfig
}

type DBConfig struct {
	PgUser     string `env:"PGUSER"`
	PgPassword string `env:"PGPASSWORD"`
	PgHost     string `env:"PGHOST"`
	PgPort     uint16 `env:"PGPORT"`
	PgDatabase string `env:"PGDATABASE"`
	PgSSLMode  string `env:"PGSSLMODE"`
}

type ServerConfig struct {
	Port        string `env:"HTTP_PORT" envDefault:"8080"`
	Timeout     string `env:"TIMEOUT" envDefault:"5s"`
	IdleTimeout string `env:"IDLE_TIMEOUT" envDefault:"30s"`
	SecretKey   string `env:"SECRET_KEY,notEmpty"`
}

type LoggerConfig struct {
	Mode     string `env:"LOG_MODE" envDefault:"info"`
	Filepath string `env:"LOG_FILE"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func New() (*Config, error) {
	config := &Config{}

	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("failed to parse config from environment variables: %w", err)
	}

	return config, nil
}
