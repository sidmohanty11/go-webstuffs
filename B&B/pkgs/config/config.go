package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/models"
)

//AppConfig -> holds the application configuration. It's like React ContextApi, stores value over all packages!
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	MailChan      chan models.MailData
}
