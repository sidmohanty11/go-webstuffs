package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/forms"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/models"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/render"
)

//is the respository type
type Repository struct {
	App *config.AppConfig
}

//the repository used by the handlers
var Repo *Repository

//creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page Handler GET "/"
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

//About is the about page handler  GET"/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

//Generals is the Generals page handler GET "/rooms/generals"
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

//Presidentals is the Presidentals page handler GET "/rooms/presidentals"
func (m *Repository) Presidentals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "presidentals.page.html", &models.TemplateData{})
}

//Reservation is the Reservation page handler GET "/make-reservation"
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})

	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

//PostReservation is the PostReservation page handler POST "/make-reservation"
//Gets the form data and checks for validity then, if True => proceeds to "/reservation-summary"
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatalln(err.Error())
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3, r)
	form.MinLength("last_name", 3, r)
	form.EmailCheck("email")
	form.PhoneCheck("phone", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

//Availability is the Search-Availability page handler GET "search-availability"
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

//PostAvailability is the PostAvailability page handler POST "search-availability"
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

//the json struct just to workout the status of availability of suites
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//AvailabilityJSON is func getting response as json
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")

	if err != nil {
		log.Fatalln(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

//Contact is the Contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

//Gets data from "/make-reservation" and then shows them here! GET "/reservation-summary"
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation summary, Maybe you should fill it first?")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
}
