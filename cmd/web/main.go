package main

import (
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/config"
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/render"
	"log"
	"net/http"
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

	render.SetApp(&app)

	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("Running app on port", portNumber)

	_ = server.ListenAndServe()
}
