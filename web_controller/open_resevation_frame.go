package web_controller

import (
	"encoding/json"
	"fmt"
	"io"
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

type openReservationFrameInput struct {
	SearchQuery   string `json:"search_query"`
	UpdateSummary string `json:"update_summary"`
}

func (c *openReservationFrameController) OpenReservationFrame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	input := &openReservationFrameInput{}

	b, err := io.ReadAll(r.Body)

	// TODO: HTTP internal error response
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error : %v", err), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, input)

	// TODO: HTTP internal error response
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error : %v", err), http.StatusInternalServerError)
		return
	}

	err = c.usecase.OpenReservationFrame(input.SearchQuery, input.UpdateSummary)

	// TODO: HTTP internal error response
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
