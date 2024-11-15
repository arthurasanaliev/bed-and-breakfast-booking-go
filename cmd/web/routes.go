package main

import (
	"github.com/arthurasanaliev/bed-and-breakfast-booking-go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes defines routing
func routes() http.Handler {
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)
	mux.Use(loadSession)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/book", handlers.Repo.Book)

	// TODO -- learn about file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
