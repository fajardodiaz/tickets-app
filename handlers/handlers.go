package handlers

import (
	"net/http"

	"github.com/fajardodiaz/tickets-app/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.page.tmpl")
}

func ListTickets(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "list.page.tmpl")
}

func NewTicket(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "new.page.tmpl")
}
