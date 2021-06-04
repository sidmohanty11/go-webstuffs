package handlers

import (
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/models"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/render"
	"net/http"
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

//Home is the home page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Passing string values from go."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
