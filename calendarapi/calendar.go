package calendarapi

import (
	"context"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"google-calendar-poc/models"
)

type CalendarService struct {
	Config *oauth2.Config
}

func NewCalendarService(cfg *oauth2.Config) *CalendarService {
	return &CalendarService{Config: cfg}
}

func (cs *CalendarService) CreateEvent(ctx context.Context, code string) (*models.CalendarEvent, error) {
	tok, err := cs.Config.Exchange(ctx, code)
	if err != nil {
		log.Printf("Token exchange error: %v", err)
		return nil, err
	}

	client := cs.Config.Client(ctx, tok)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Calendar service creation error: %v", err)
		return nil, err
	}

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Printf("Timezone error: %v", err)
		return nil, err
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
		return nil, err
	}

	return models.FromGoogleEvent(createdEvent), nil
}
