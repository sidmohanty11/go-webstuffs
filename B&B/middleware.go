package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//adds CSRF -> cross site request forgery protection to all POST reqs
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//loads and saves the session on every req
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
