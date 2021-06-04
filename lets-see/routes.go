package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sidmohanty11/go-webstuffs/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/pkgs/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
