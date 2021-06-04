package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sidmohanty11/go-webstuffs/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/pkgs/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//for serving static files =>
	// fileServer := http.FileServer(http.Dir("./static/"))
	// mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
