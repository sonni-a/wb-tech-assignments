package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func isForm(r *http.Request) bool {
	return strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded")
}

func parseCreateRequest(r *http.Request) (userID int, date, text string, err error) {
	if isForm(r) {
		if err := r.ParseForm(); err != nil {
			return 0, "", "", err
		}
		userID, _ = strconv.Atoi(r.FormValue("user_id"))
		return userID, r.FormValue("date"), r.FormValue("event"), nil
	}

	var req struct {
		UserID int    `json:"user_id"`
		Date   string `json:"date"`
		Text   string `json:"event"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return 0, "", "", err
	}
	return req.UserID, req.Date, req.Text, nil
}

func parseUpdateRequest(r *http.Request) (id int, text string, err error) {
	if isForm(r) {
		if err := r.ParseForm(); err != nil {
			return 0, "", err
		}
		id, _ = strconv.Atoi(r.FormValue("id"))
		return id, r.FormValue("event"), nil
	}

	var req struct {
		ID   int    `json:"id"`
		Text string `json:"event"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return 0, "", err
	}
	return req.ID, req.Text, nil
}

func parseDeleteRequest(r *http.Request) (id int, err error) {
	if isForm(r) {
		if err := r.ParseForm(); err != nil {
			return 0, err
		}
		id, _ = strconv.Atoi(r.FormValue("id"))
		return id, nil
	}

	var req struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return 0, err
	}
	return req.ID, nil
}
