package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"calendar/internal/storage"
)

type Handler struct {
	cal *storage.Calendar
}

func NewHandler(cal *storage.Calendar) *Handler {
	return &Handler{cal: cal}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		_, _ = w.Write([]byte(`{"error":"internal error"}`))
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func requireMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return false
	}
	return true
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodPost) {
		return
	}

	userID, dateStr, text, err := parseCreateRequest(r)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid request"})
		return
	}

	date, err := parseDate(dateStr)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}

	event, _ := h.cal.Create(userID, date, text)

	writeJSON(w, 200, map[string]any{"result": event})
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodPost) {
		return
	}

	id, text, err := parseUpdateRequest(r)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid request"})
		return
	}

	if err := h.cal.Update(id, text); err != nil {
		writeJSON(w, 503, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, 200, map[string]string{"result": "updated"})
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodPost) {
		return
	}

	id, err := parseDeleteRequest(r)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid request"})
		return
	}

	if err := h.cal.Delete(id); err != nil {
		writeJSON(w, 503, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, 200, map[string]string{"result": "deleted"})
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodGet) {
		return
	}

	h.eventsForPeriod(w, r, "day")
}

func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodGet) {
		return
	}

	h.eventsForPeriod(w, r, "week")
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodGet) {
		return
	}

	h.eventsForPeriod(w, r, "month")
}

func (h *Handler) eventsForPeriod(w http.ResponseWriter, r *http.Request, mode string) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	date, err := parseDate(r.URL.Query().Get("date"))
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}

	var from, to time.Time

	switch mode {
	case "day":
		from = date
		to = date
	case "week":
		wd := int(date.Weekday())
		if wd == 0 {
			wd = 7
		}
		from = date.AddDate(0, 0, -(wd-1))
		to = from.AddDate(0, 0, 6)
	case "month":
		from = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
		to = from.AddDate(0, 1, -1)
	}

	events := h.cal.EventsForPeriod(userID, from, to)
	writeJSON(w, 200, map[string]any{"result": events})
}
