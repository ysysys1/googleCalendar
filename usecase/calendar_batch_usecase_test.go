package usecase

import (
	"errors"
	"testing"

	mock_outbounds "github.com/calendar-open/outbounds/mock"
	"github.com/golang/mock/gomock"
	"google.golang.org/api/calendar/v3"
)

func TestOpenCalendarUsecase_OpenReservationFrame(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		mockFn  func(*mock_outbounds.MockCalendarClient)
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func(mcc *mock_outbounds.MockCalendarClient) {
				events := &calendar.Events{
					Items: []*calendar.Event{
						{
							Id:      "Id",
							Summary: "a",
						},
					},
				}
				mcc.EXPECT().SearchEvents(gomock.Any()).Return(events, nil)
				mcc.EXPECT().UpdateEventSummary(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			wantErr: false,
		},
		{
			name: "when unable to fetch events",
			mockFn: func(mcc *mock_outbounds.MockCalendarClient) {
				mcc.EXPECT().SearchEvents(gomock.Any()).Return(nil, errors.New("Some error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mock_outbounds.NewMockCalendarClient(ctrl)
			tt.mockFn(m)
			u := &calendarBatchUsecase{
				CalendarClient: m,
			}
			if err := u.OpenReservationFrame(); (err != nil) != tt.wantErr {
				t.Errorf("OpenCalendarUsecase.OpenReservationFrame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
