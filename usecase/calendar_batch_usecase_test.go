package usecase

import (
	"errors"
	"testing"

	mock_environments "github.com/calendar-open/environments/mock"
	mock_outbounds "github.com/calendar-open/outbounds/mock"
	"github.com/golang/mock/gomock"
	"google.golang.org/api/calendar/v3"
)

func TestOpenCalendarUsecase_OpenReservationFrame(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		searchQuery   string
		updateSummary string
	}

	tests := []struct {
		name      string
		mockFn    func(*mock_outbounds.MockCalendarClient)
		envMockFn func(*mock_environments.MockEnvironments)
		args      args
		wantErr   bool
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
			envMockFn: func(me *mock_environments.MockEnvironments) {
				me.EXPECT().GetInvalidWordsSearchQueryString().Return(nil)
			},
			args: args{
				searchQuery:   "test",
				updateSummary: "dummy",
			},
			wantErr: false,
		},
		{
			name: "when unable to fetch events",
			mockFn: func(mcc *mock_outbounds.MockCalendarClient) {
				mcc.EXPECT().SearchEvents(gomock.Any()).Return(nil, errors.New("Some error"))
			},
			envMockFn: func(me *mock_environments.MockEnvironments) {
				me.EXPECT().GetInvalidWordsSearchQueryString().Return(nil)
			},
			args: args{
				searchQuery:   "test",
				updateSummary: "dummy",
			},
			wantErr: true,
		},
		{
			name:   "when invalid search query",
			mockFn: func(mcc *mock_outbounds.MockCalendarClient) {},
			envMockFn: func(me *mock_environments.MockEnvironments) {
				me.EXPECT().GetInvalidWordsSearchQueryString().Return([]string{"test"})
			},
			args: args{
				searchQuery:   "test",
				updateSummary: "dummy",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mock_outbounds.NewMockCalendarClient(ctrl)
			me := mock_environments.NewMockEnvironments(ctrl)
			tt.mockFn(m)
			tt.envMockFn(me)
			u := &calendarBatchUsecase{
				CalendarClient: m,
				Env:            me,
			}
			if err := u.OpenReservationFrame(tt.args.searchQuery, tt.args.updateSummary); (err != nil) != tt.wantErr {
				t.Errorf("OpenCalendarUsecase.OpenReservationFrame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
