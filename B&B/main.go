package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/driver"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/handlers"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/helpers"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/models"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/render"
)

//The port you want to serve to
const PORT = ":8000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	//-------put it back when testing is done ------
	//put stuff in session =>
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	//read flags
	inProduction := flag.Bool("production", true, "App in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbHost := flag.String("dbhost", "localhost", "database host")
	dbName := flag.String("dbname", "", "database name")
	dbUser := flag.String("dbuser", "", "database user")
	dbPassword := flag.String("dbpassword", "", "database password")
	dbPort := flag.String("dbport", "5432", "database port")
	dbSSL := flag.String("dbssl", "disable", "database ssl settings (disable, prefer,require)")

	flag.Parse()

	if *dbName == "" || *dbUser == "" || *dbPassword == "" {
		fmt.Println("Missing required flags")
	}

	app.MailChan = make(chan models.MailData)
	defer close(app.MailChan) //put in app.MailChan when testing so main can execute fully
	listenForMail()

	// format of email
	fmt.Println("Starting mail listener!")

	//url -> uniform resource locator!
	//true when in prod
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	//connect to DB!
	log.Println("Connecting to DB at PORT 5432")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPassword, *dbSSL)
	db, err := driver.ConnectSQL(connectionString)

	if err != nil {
		log.Fatalln("Cannot connect to DB!")
	}

	log.Println("Connected to DB at PORT 5432")
	defer db.SQL.Close() //if testing -> run func returns db too because after run() the db will close if this is there

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Println(err.Error())
		return
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	//---------------------------------------------
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

// func run() error {
// 	return nil
// }
