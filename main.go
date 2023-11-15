package main

import (
	"net/http"

	"github.com/fajardodiaz/tickets-app/handlers"
)

func main() {
	PortNumber := ":8081"
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/tickets", handlers.ListTickets)
	http.HandleFunc("/new", handlers.NewTicket)

	http.ListenAndServe(PortNumber, nil)
}
