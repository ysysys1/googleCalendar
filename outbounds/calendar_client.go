package outbounds

import (
	"context"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var calendarID string = os.Getenv("GOOGLE_CALENDAR_ID")

type CalendarClient interface {
	SearchEvents(serachQuery string) (*calendar.Events, error)
	UpdateEventSummary(targetEvent *calendar.Event, summary string) (*calendar.Event, error)
}

type calendarClient struct {
	ClientOption option.ClientOption
}

func NewCalendarClient() CalendarClient {
	return &calendarClient{
		ClientOption: client(),
	}
}

func client() option.ClientOption {
	ctx := context.Background()
	scope := []string{calendar.CalendarScope, calendar.CalendarEventsScope}

	key := os.Getenv("GOOGLE_API_CREDENTIAL")
	replacedKey := strings.Replace(key, "\n", "\\n", -1)

	config, err := google.JWTConfigFromJSON([]byte(replacedKey), scope...)
	if err != nil {
		log.Printf("Unable to init jwt config: %v", err)
	}
	tokenSource := config.TokenSource(ctx)
	return option.WithTokenSource(tokenSource)
}

func (c *calendarClient) SearchEvents(serachQuery string) (*calendar.Events, error) {
	ctx := context.Background()

	srv, err := calendar.NewService(ctx, c.ClientOption)
	if err != nil {
		log.Printf("Unable to retrieve Calendar client: %v", err)
		return nil, err
	}
	list := srv.Events.List(calendarID)
	list.MaxResults(999)
	list.Q(serachQuery)
	return list.Do()
}

func (c *calendarClient) UpdateEventSummary(targetEvent *calendar.Event, summary string) (*calendar.Event, error) {
	ctx := context.Background()

	srv, err := calendar.NewService(ctx, c.ClientOption)
	if err != nil {
		log.Printf("Unable to retrieve Calendar client: %v", err)
		return nil, err
	}

	targetEvent.Summary = summary

	eventCall := srv.Events.Update(calendarID, targetEvent.Id, targetEvent)
	return eventCall.Do()
}
