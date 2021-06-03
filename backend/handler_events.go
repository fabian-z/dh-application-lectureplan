package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		panic(err)
	}
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
