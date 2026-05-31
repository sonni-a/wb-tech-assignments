package main

import (
	"log"
	"net/http"
	"os"

	apphttp "calendar/internal/http"
	"calendar/internal/storage"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cal := storage.NewCalendar()
	h := apphttp.NewHandler(cal)

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.EventsForDay)
	mux.HandleFunc("/events_for_week", h.EventsForWeek)
	mux.HandleFunc("/events_for_month", h.EventsForMonth)

	log.Println("server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, apphttp.Logging(mux)))
}
