package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type EventResponse struct {
	ID int64 `json:"id"`
	//GroupID not useful here
	AllDay bool      `json:"allDay"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"` //Exlusive value!
	Title  string    `json:"title"`
}

// Basic query all events
func handleListEvents(w http.ResponseWriter, r *http.Request) {
	tx, err := db.Beginx()
	if err != nil {
		panic(err)
	}
	events, err := GetEvents(tx, "SELECT * FROM events")
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	event := events[0]
	startTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), int(event.StartTime.Hour), int(event.StartTime.Minute), int(event.StartTime.Second), 0, time.Local)
	endTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), int(event.EndTime.Hour), int(event.EndTime.Minute), int(event.EndTime.Second), 0, time.Local)

	response := EventResponse{
		ID:    events[0].ID,
		Start: startTime,
		End:   endTime,
		Title: event.Lecture.Name,
	}

	err = json.NewEncoder(w).Encode([]EventResponse{response})
	if err != nil {
		panic(err)
	}
}

// Basic endpoint
func handleEvents(w http.ResponseWriter, r *http.Request) {
	// TODO initial session creation
	err := templates.Events.Execute(w, struct {
		Title       string
		PageTitle   string
		ShowActions bool
	}{
		"DHBW Lörrach - Vorlesungsplanung",
		"Termine",
		true,
	})
	if err != nil {
		log.Println("Request error: ", err)
	}
}
