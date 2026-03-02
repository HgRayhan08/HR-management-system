package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	fmt.Println("sudah menemukan .env")

	expString, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Tz:       os.Getenv("DB_TZ"),
		},
		JWT: JWT{
			Secret: os.Getenv("JWT_KEY"),
			Exp:    expString,
		},
	}
}
