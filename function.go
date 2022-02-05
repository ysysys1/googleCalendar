package function

import (
	"net/http"

	"github.com/calendar-open/environments"
	"github.com/calendar-open/outbounds"
	"github.com/calendar-open/usecase"
	"github.com/calendar-open/web_controller"
)

func OpenCalendar(w http.ResponseWriter, r *http.Request) {
	calendarClient := outbounds.NewCalendarClient()
	envs := environments.NewEnvironments()
	usecase := usecase.NewCalendarBatchUsecase(calendarClient, envs)
	controller := web_controller.NewOpenReservationFrameWebController(usecase)
	controller.OpenReservationFrame(w, r)
}
