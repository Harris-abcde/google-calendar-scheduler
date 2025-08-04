package controllers

import (
	"context"
	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"google-calendar-poc/models"
)

type CalendarController struct {
	Config *oauth2.Config
}

func NewCalendarController(cfg *oauth2.Config) *CalendarController {
	return &CalendarController{Config: cfg}
}

func (cc *CalendarController) AuthHandler(c *gin.Context) {
	authURL := cc.Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (cc *CalendarController) CallbackHandler(c *gin.Context) {
	 log.Printf("Received callback with query params: %v", c.Request.URL.Query())

    code := c.Query("code")
    if code == "" {
        log.Println("No code parameter received in callback")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authorization code"})
        return
    }
	tok, err := cc.Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Token exchange error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	client := cc.Config.Client(context.Background(), tok)
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Calendar service creation error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create calendar service"})
		return
	}

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Printf("Timezone error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load timezone"})
		return
	}

	now := time.Now().In(loc)
	eventModel := models.NewCalendarEvent(
		"Automated Google Meet Event",
		"This event was automatically scheduled after calendar connection",
		now.Add(1*time.Hour),
		30*time.Minute,
	)

	createdEvent, err := srv.Events.Insert("primary", eventModel.ToGoogleEvent()).ConferenceDataVersion(1).Do()
	if err != nil {
		log.Printf("Event creation error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create calendar event"})
		return
	}

	c.JSON(http.StatusOK, models.FromGoogleEvent(createdEvent))
}