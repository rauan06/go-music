package config

import (
	"go-music/models"
	"go-music/pkg/config/flags"
	"go-music/pkg/config/logger"
	"log"
	"os"
)

var Config *models.Config

// SetupConfig parses command-line arguments, environment variables from a YAML file, and configures the logger.
// If the help flag is provided, the program prints the usage information and exits with code 0.
// If an error occurs during flag validation, the program terminates with code 1.
func SetupConfig() *models.Config {
	conf := newConfig()

	if err := flags.ParseFlags(conf); err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	logger.SetLogger(conf)

	conf.Username = os.Getenv("USERNAME")
	conf.Email = os.Getenv("EMAIL")
	conf.Pass = os.Getenv("PASSWORD")

	Config = conf
	return conf
}

func newConfig() *models.Config {
	return &models.Config{
		Port: 8000,
		Env:  "prod",
	}
}
