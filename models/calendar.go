package models

import (
	"fmt"
	"time"

	"google.golang.org/api/calendar/v3"
)

type CalendarEvent struct {
	Summary      string        `json:"summary"`
	Description  string        `json:"description"`
	Start        EventDateTime `json:"start"`
	End          EventDateTime `json:"end"`
	MeetLink     string        `json:"meet_link,omitempty"`
	CalendarLink string        `json:"calendar_link,omitempty"`
}

type EventDateTime struct {
	DateTime time.Time `json:"date_time"`
	TimeZone string    `json:"time_zone"`
}

func (ce *CalendarEvent) ToGoogleEvent() *calendar.Event {
	return &calendar.Event{
		Summary:     ce.Summary,
		Description: ce.Description,
		Start: &calendar.EventDateTime{
			DateTime: ce.Start.DateTime.Format(time.RFC3339),
			TimeZone: ce.Start.TimeZone,
		},
		End: &calendar.EventDateTime{
			DateTime: ce.End.DateTime.Format(time.RFC3339),
			TimeZone: ce.End.TimeZone,
		},
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: fmt.Sprintf("meet-%d", time.Now().Unix()),
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
			},
		},
	}
}

func FromGoogleEvent(ge *calendar.Event) *CalendarEvent {
	startTime, _ := time.Parse(time.RFC3339, ge.Start.DateTime)
	endTime, _ := time.Parse(time.RFC3339, ge.End.DateTime)

	var meetLink string
	if ge.ConferenceData != nil && len(ge.ConferenceData.EntryPoints) > 0 {
		meetLink = ge.ConferenceData.EntryPoints[0].Uri
	}

	return &CalendarEvent{
		Summary:      ge.Summary,
		Description:  ge.Description,
		Start:        EventDateTime{DateTime: startTime, TimeZone: ge.Start.TimeZone},
		End:          EventDateTime{DateTime: endTime, TimeZone: ge.End.TimeZone},
		MeetLink:     meetLink,
		CalendarLink: ge.HtmlLink,
	}
}

func NewCalendarEvent(summary, description string, startTime time.Time, duration time.Duration) *CalendarEvent {
	return &CalendarEvent{
		Summary:     summary,
		Description: description,
		Start: EventDateTime{
			DateTime: startTime,
			TimeZone: "Asia/Kolkata",
		},
		End: EventDateTime{
			DateTime: startTime.Add(duration),
			TimeZone: "Asia/Kolkata",
		},
	}
}