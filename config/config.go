package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret   string
	FrontendURL string
	Port        string
	PostgresURL string
}

// Checks that none of the .env variables are empty
func (config *Config) IsValid() bool {
	return len(config.JWTSecret) != 0 &&
		len(config.FrontendURL) != 0 &&
		len(config.Port) != 0
}

var Vars *Config

// Initialize loads environment variables from the .env file,
// including JWT_SECRET, FRONTEND_URL, PORT, and POSTGRES_URL.
// It panics if any required config is missing or invalid.
func Initialize() {
	godotenv.Load(".env")

	Vars = new(Config)
	Vars.JWTSecret = os.Getenv("JWT_SECRET")
	Vars.FrontendURL = os.Getenv("FRONTEND_URL")
	Vars.Port = os.Getenv("PORT")
	Vars.PostgresURL = os.Getenv("POSTGRES_URL")

	if !Vars.IsValid() {
		panic("failed to initialize config")
	}
}
