package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	mux.Get("/rooms/generals", handlers.Repo.Generals)
	mux.Get("/rooms/presidentals", handlers.Repo.Presidentals)

	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/")) //for serving static files
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
