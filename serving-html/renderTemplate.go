package main

import (
	"html/template"
	"log"
	"net/http"
)

//renderTemplate func -> executes the templates
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
