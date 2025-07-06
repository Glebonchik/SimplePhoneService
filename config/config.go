package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataBase   DataBaseConfig
	HTTPServer HTTPServer
}

type DataBaseConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

type HTTPServer struct {
	Port string
}

func LoadConfig() (*Config, error) {
	// Только для dev-сервера
	if err := godotenv.Load(); err != nil {
		log.Println("No .env warning file found, using environment variables")
	}

	cfg := &Config{}

	cfg.DataBase.DBHost = os.Getenv("DB_HOST")
	if cfg.DataBase.DBHost == "" {
		return nil, fmt.Errorf("DB_HOST is required")
	}

	cfg.DataBase.DBPort = os.Getenv("DB_PORT")
	if cfg.DataBase.DBPort == "" {
		cfg.DataBase.DBPort = ""
	}

	cfg.DataBase.DBUser = os.Getenv("DB_USER")
	if cfg.DataBase.DBUser == "" {
		return nil, fmt.Errorf("DB_USER is required")
	}

	cfg.DataBase.DBPassword = os.Getenv("DB_PASSWORD")
	if cfg.DataBase.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	cfg.DataBase.DBName = os.Getenv("DB_NAME")
	if cfg.DataBase.DBName == "" {
		return nil, fmt.Errorf("DB_NAME is required")
	}

	cfg.DataBase.DBSSLMode = os.Getenv("DB_SSLMODE")
	if cfg.DataBase.DBSSLMode == "" {
		cfg.DataBase.DBSSLMode = "disable"
	}

	cfg.HTTPServer.Port = os.Getenv("SERVER_PORT")
	if cfg.HTTPServer.Port == "" {
		cfg.HTTPServer.Port = ":8082"
	}

	return cfg, nil
}
