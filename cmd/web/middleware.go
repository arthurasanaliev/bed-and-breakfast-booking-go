package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// noSurf adds CSRF protection to all POST requests
func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// loadSession loads and saves the session on every request
func loadSession(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
