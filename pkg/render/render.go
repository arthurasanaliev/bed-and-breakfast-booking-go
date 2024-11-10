package render

import (
	"bytes"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/config"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// SetApp sets app passed from main package
func SetApp(a *config.AppConfig) {
	app = a
}

// addDefaultData adds default data
func addDefaultData(tmplData *models.TemplateData) *models.TemplateData {
    // TODO -- add when needed
    return tmplData
}

// RenderTemplate renders template from cache
func RenderTemplate(w http.ResponseWriter, t string, tmplData *models.TemplateData) {
	var cache map[string]*template.Template
	if app.UseCache {
		log.Println("used cache for", t)
		cache = app.TemplateCache
	} else {
		log.Println("created cache for", t)
		cache, _ = CreateTemplateCache()
	}

	tmpl, ok := cache[t]
	if !ok {
		log.Fatal("cannot get template from cache")
	}

    tmplData = addDefaultData(tmplData)

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, tmplData)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache creates template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*")
	if err != nil {
		return cache, err
	}

	layouts, err := filepath.Glob("./templates/layouts/*")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		set, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			set, err = set.ParseGlob("./templates/layouts/*")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = set
	}

	return cache, nil
}
