package handlers

import (
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/models"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/render"
	"net/http"
)

// Home is the home-page handler function
func Home(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello from template data!"

	render.RenderTemplate(w, "home.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
