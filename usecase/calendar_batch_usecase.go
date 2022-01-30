package usecase

import (
	"log"
	"time"

	"github.com/calendar-open/outbounds"
)

type CalendarBatchUsecase interface {
	OpenReservationFrame() error
}

type calendarBatchUsecase struct {
	CalendarClient outbounds.CalendarClient
}

func NewCalendarBatchUsecase(client outbounds.CalendarClient) CalendarBatchUsecase {
	return &calendarBatchUsecase{
		CalendarClient: client,
	}
}

func (u *calendarBatchUsecase) OpenReservationFrame() error {
	events, err := u.CalendarClient.SearchEvents("準備枠")

	if err != nil {
		log.Printf("Unable to fetch events: %v", err)
		return err
	}

	for _, event := range events.Items {
		time.Sleep(time.Second * 1)
		_, err := u.CalendarClient.UpdateEventSummary(event, "予約枠")

		if err != nil {
			return err
		}
	}

	return nil
}
