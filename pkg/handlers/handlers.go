package handlers

import (
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/config"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/models"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/render"
	"net/http"
)

// Repository stores data for handlers package
type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repository instance
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

// SetRepo sets Repository instance
func SetRepo(r *Repository) {
	Repo = r
}

// Home is the home-page handler function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	visited := "Home was visited"
	m.app.Session.Put(r.Context(), "visited", visited)

	render.RenderTemplate(w, "home.tmpl", &models.TemplateData{})
}

// Book is the booking-page handler function
func (m *Repository) Book(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}

	visited := m.app.Session.GetString(r.Context(), "visited")
	stringMap["visited"] = visited

	render.RenderTemplate(w, "book.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
