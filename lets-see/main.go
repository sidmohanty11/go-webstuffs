package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sidmohanty11/go-webstuffs/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/pkgs/handlers"
	"github.com/sidmohanty11/go-webstuffs/pkgs/render"
)

//The port you want to serve to
const PORT = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//url -> uniform resource locator!
	//true when in prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalln(err.Error())
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Listening at PORT%s", PORT)
	fmt.Println()

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
