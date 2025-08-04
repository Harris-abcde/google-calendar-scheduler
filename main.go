package main

import (
	"google-calendar-poc/config"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	// "google-calendar-gin/config"
	"google-calendar-poc/routes"
)

func main() {
	appConfig := config.LoadConfig()

	oauthConfig := &oauth2.Config{
		ClientID:     appConfig.GoogleClientID,
		ClientSecret: appConfig.GoogleClientSecret,
		RedirectURL:  appConfig.RedirectURL,
		Scopes:       []string{appConfig.CalendarScope},
		Endpoint:     google.Endpoint,
	}

	router := gin.Default()
	routes.SetupRoutes(router, oauthConfig)

	port := appConfig.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}