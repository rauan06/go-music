package flags

import (
	"flag"
	"fmt"
	"go-music/models"
	"os"
)

// ParseFlags parses command-line flags and validates the configuration.
func ParseFlags(config *models.Config) error {
	flag.BoolVar(&config.Help, "help", false, "prints usage information")
	flag.IntVar(&config.Port, "port", 8000, "sets the port number (1024-49151)")
	flag.StringVar(&config.Env, "env", "prod", "sets the environment ('local', 'dev', 'prod')")

	flag.Parse()

	if err := validateConfig(config); err != nil {
		return err
	}

	return nil
}

// validateConfig validates the configuration values.
func validateConfig(config *models.Config) error {
	if config.Help {
		printUsage()
		os.Exit(0)
	}

	if config.Port < 1024 || config.Port > 49151 {
		return fmt.Errorf("invalid port number: %d, accepted range is 1024 - 49151", config.Port)
	}

	if config.Env != "local" && config.Env != "dev" && config.Env != "prod" {
		return fmt.Errorf("invalid environment: %s, accepted values are: 'local', 'dev', 'prod'", config.Env)
	}

	return nil
}

func printUsage() {
	fmt.Println(`Go music website

Usage:
  ./go-music [--port <N>] [--env <S>]
  ./go-music --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --env S      Environment variable ('local', 'dev', 'prod').`)
}
