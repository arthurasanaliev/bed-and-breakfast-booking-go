package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/config"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/handlers"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig

// main is the entry point of the program
func main() {
	app.UseCache = true
	app.InProduction = false

	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = cache

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	render.SetApp(&app)

	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("Running app on port", portNumber)

	_ = server.ListenAndServe()
}
