package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//NewRenderer -> sets the config for the template pkg
func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	if app.Session.Exists(r.Context(), "user_id") {
		td.IsAuthenticated = 1
	}
	return td
}

//Template func -> executes the templates
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache { // in production -> cache it!
		tc = app.TemplateCache
	} else { //in development mode - read from disk
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]

	if !ok {
		log.Fatalln("Couldn't get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
}

//CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html") //finding pages.html

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	for _, page := range pages {
		name := filepath.Base(page)                                     //name of the page html found
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) //parsing tmpl

		if err != nil {
			log.Fatalln("Error:", err.Error())
		}

		matches, err := filepath.Glob("./templates/*.layout.html") //finding layout.html

		if err != nil {
			log.Fatalln("Error:", err.Error())
		}

		if len(matches) > 0 { //if there are layouts
			ts, err = ts.ParseGlob("./templates/*.layout.html") //parse layouts

			if err != nil {
				log.Fatalln("Error:", err.Error())
			}
		}

		cache[name] = ts //store key=>.page.html, value=>.layout.html
	}
	return cache, nil //returns the tmpl pointer map
}
