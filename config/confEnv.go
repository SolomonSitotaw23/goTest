package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GIN_MODE   string
	PORT       string
	DBHOST     string
	DBPORT     string
	DBNAME     string
	DBUSER     string
	DBPASSWORD string
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return &Config{
		GIN_MODE:   os.Getenv("GIN_MODE"),
		PORT:       os.Getenv("PORT"),
		DBHOST:     os.Getenv("DBHOST"),
		DBPORT:     os.Getenv("DBPORT"),
		DBNAME:     os.Getenv("DBNAME"),
		DBUSER:     os.Getenv("DBUSER"),
		DBPASSWORD: os.Getenv("DBPASSWORD"),
	}
}
