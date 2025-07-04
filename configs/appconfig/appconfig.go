package appconfig

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string // Port for the web server
	DbHost    string // Database host
	DbUser    string // Database user
	DbPass    string // Database password
	DbName    string // Database name
	DbPort    string // Database port
	DbCharset string // Database charset
	DbLoc     string // Database location
}

func LoadConfig() *Config {
	_ = godotenv.Load() // Load environment variables from .env file

	config := &Config{
		Port:      Getenv("PORT", "8080"),
		DbHost:    Getenv("DB_HOST", "localhost"),
		DbUser:    Getenv("DB_USER", "admin"),
		DbPass:    Getenv("DB_PASS", "admin123"),
		DbName:    Getenv("DB_NAME", "mywebapp"),
		DbPort:    Getenv("DB_PORT", "3306"),
		DbCharset: Getenv("DB_CHARSET", "utf8mb4"),
		DbLoc:     Getenv("DB_LOC", "Asia%2FShanghai"),
	}

	return config
}

func Getenv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
