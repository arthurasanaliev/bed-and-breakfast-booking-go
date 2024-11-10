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

	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)

	mux.Get("/", handlers.Home)

	return mux
}
