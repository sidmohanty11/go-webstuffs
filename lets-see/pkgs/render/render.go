package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

//RenderTemplate func -> executes the templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache() //tc:templateCache from CreateTemplateCache func

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatalln(err.Error())
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)

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
