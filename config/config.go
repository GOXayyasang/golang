package config

import (
	"net/url"
	"os"
)

// DatabaseConfig contains the database connection configuration
type DatabaseConfig struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetDatabaseConfig returns the database configuration
func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Server:   os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"), // Default MSSQL port
		User:     os.Getenv("DB_USER"),
		Password: url.QueryEscape(os.Getenv("DB_PASSWORD")),
		Database: os.Getenv("DB_NAME"),
	}
}
