package web_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/calendar-open/usecase"
)

type OpenReservationFrameWebController interface {
	OpenReservationFrame(w http.ResponseWriter, r *http.Request)
}

type openReservationFrameController struct {
	usecase usecase.CalendarBatchUsecase
}

func NewOpenReservationFrameWebController(usecase usecase.CalendarBatchUsecase) OpenReservationFrameWebController {
	return &openReservationFrameController{
		usecase: usecase,
	}
}

func (c *openReservationFrameController) OpenReservationFrame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := c.usecase.OpenReservationFrame()

	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error : %v", err), http.StatusInternalServerError)
		return
	}

	res := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error : %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
