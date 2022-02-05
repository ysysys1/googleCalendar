package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/calendar-open/environments"
	"github.com/calendar-open/outbounds"
)

type CalendarBatchUsecase interface {
	OpenReservationFrame(searchQuery string, updateSummary string) error
}

type calendarBatchUsecase struct {
	CalendarClient outbounds.CalendarClient
	Env            environments.Environments
}

func NewCalendarBatchUsecase(client outbounds.CalendarClient, env environments.Environments) CalendarBatchUsecase {
	return &calendarBatchUsecase{
		CalendarClient: client,
		Env:            env,
	}
}

func (u *calendarBatchUsecase) OpenReservationFrame(searchQuery string, updateSummary string) error {
	if valid := u.isValidateSearchQuery(searchQuery); !valid {
		return fmt.Errorf("invalid Update Query :%s", searchQuery)
	}

	events, err := u.CalendarClient.SearchEvents(searchQuery)

	if err != nil {
		log.Printf("Unable to fetch events: %v", err)
		return err
	}

	for _, event := range events.Items {
		time.Sleep(time.Second * 1)
		_, err := u.CalendarClient.UpdateEventSummary(event, updateSummary)

		if err != nil {
			return err
		}
	}

	return nil
}

func (u *calendarBatchUsecase) isValidateSearchQuery(searchQuery string) bool {
	words := u.Env.GetInvalidWordsSearchQueryString()

	for _, word := range words {
		if word == searchQuery {
			return false
		}
	}
	return true
}
