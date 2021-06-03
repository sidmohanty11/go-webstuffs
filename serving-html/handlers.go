package main

import (
	"net/http"
)

//Home is the home page Handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
