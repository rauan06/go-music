package logger

import (
	"go-music/models"
	"log/slog"
	"os"
)

const (
	envTest = "test"
	envProd = "prod"
)

// SetLogger updates the Logger value in the configuration based on the specified Env.
func SetLogger(conf *models.Config) {
	switch conf.Env {
	case envTest:
		conf.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		conf.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
}
