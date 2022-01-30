package function

import (
	"net/http"

	"github.com/calendar-open/outbounds"
	"github.com/calendar-open/usecase"
	"github.com/calendar-open/web_controller"
)

func OpenCalendar(w http.ResponseWriter, r *http.Request) {
	calendarClient := outbounds.NewCalendarClient()
	usecase := usecase.NewCalendarBatchUsecase(calendarClient)
	controller := web_controller.NewOpenReservationFrameWebController(usecase)
	controller.OpenReservationFrame(w, r)
}
