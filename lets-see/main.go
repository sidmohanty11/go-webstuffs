package main

import (
	"fmt"
	"github.com/sidmohanty11/go-webstuffs/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/pkgs/handlers"
	"github.com/sidmohanty11/go-webstuffs/pkgs/render"
	"log"
	"net/http"
)

//The port you want to serve to
const PORT = ":8000"

func main() {
	//url -> uniform resource locator!
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalln(err.Error())
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Listening at PORT%s", PORT)
	fmt.Println()
	http.ListenAndServe(PORT, nil)
}
