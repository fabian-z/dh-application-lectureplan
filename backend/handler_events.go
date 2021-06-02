package main

import (
	"log"
	"net/http"
)

// Basic query all events
func handleListEvents(w http.ResponseWriter, r *http.Request) {
	tx, err := db.Beginx()
	if err != nil {
		panic(err)
	}
	var events []Event
	tx.Select(&events, "SELECT * FROM events")
}

// Basic endpoint
func handleEvents(w http.ResponseWriter, r *http.Request) {
	// TODO initial session creation
	err := templates.Events.Execute(w, struct{ Title string }{
		"DHBW Lörrach - Vorlesungsplanung",
	})
	if err != nil {
		log.Println("Request error: ", err)
	}
}
