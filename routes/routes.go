package routes

import (
	"google-calendar-poc/controllers"

	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
)

func SetupRoutes(router *gin.Engine, cfg *oauth2.Config) {
	calendarController := controllers.NewCalendarController(cfg)

	router.GET("/auth/google", calendarController.AuthHandler)
	router.GET("/auth/google/callback", calendarController.CallbackHandler)
}