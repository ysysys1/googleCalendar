package function

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/calendar-open/outbounds"
	"github.com/calendar-open/usecase"
)

func OpenCalendar(w http.ResponseWriter, r *http.Request) {
	calendarClient := outbounds.NewCalendarClient()
	usecase := usecase.NewCalendarBatchUsecase(calendarClient)

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	err := usecase.OpenReservationFrame()
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error : %v", err), http.StatusInternalServerError)
	}

	res := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error : %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
