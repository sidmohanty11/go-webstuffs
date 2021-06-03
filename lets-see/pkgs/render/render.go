package render

import (
	"html/template"
	"log"
	"net/http"
)

//RenderTemplate func -> executes the templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
