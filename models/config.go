package models

import (
	"fmt"
	"log/slog"
)

type Config struct {
	Help     bool
	Env      string // Default is "prod"
	Port     int    // Default is 8000
	Logger   *slog.Logger
	Username string `env:"USERNAME"` // os.Getenv("USERNAME")
	Email    string `env:"EMAIL"`    // os.Getenv("EMAIL")
	Pass     string `env:"PASSWORD"` // os.Getenv("PASSWORD")
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("0.0.0.0%d", c.Port)
}
