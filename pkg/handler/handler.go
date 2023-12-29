package handler

import (
	"net/http"

	"github.com.br/Leodf/bookings/pkg/config"
	"github.com.br/Leodf/bookings/pkg/model"
	"github.com.br/Leodf/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.Template(w, "home.page.tmpl", &model.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.Template(w, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the make reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "make-reservation.page.tmpl", &model.TemplateData{})
}

// Generals is the generals quarters page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "generals.page.tmpl", &model.TemplateData{})
}

// Majors is the majors suite page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "majors.page.tmpl", &model.TemplateData{})
}

// Majors is the majors suite page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "search-availability.page.tmpl", &model.TemplateData{})
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "contact.page.tmpl", &model.TemplateData{})
}
