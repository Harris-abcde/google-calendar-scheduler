package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/calendar/v3"
)

type Config struct {
	GoogleClientID     string
	GoogleClientSecret string
	RedirectURL        string
	Port               string
	CalendarScope      string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	return &Config{
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:        os.Getenv("REDIRECT_URL"),
		Port:               os.Getenv("PORT"),
		CalendarScope:      calendar.CalendarScope,
	}
}