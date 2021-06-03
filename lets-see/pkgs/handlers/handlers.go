package handlers

import (
	"github.com/sidmohanty11/go-webstuffs/pkgs/render"
	"net/http"
)

//Home is the home page Handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
