package main

import (
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/events", http.StatusTemporaryRedirect)
}

func handleDataProtection(w http.ResponseWriter, r *http.Request) {
	err := templates.DataProtection.Execute(w, struct {
		Title       string
		PageTitle   string
		ShowActions bool
	}{
		"DHBW Lörrach - Vorlesungsplanung",
		"Datenschutz",
		false,
	})
	if err != nil {
		log.Println("Request error: ", err)
	}
}

func emptyHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
