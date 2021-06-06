package models

import (
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/forms"
)

//holds data sent from handlers -> templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
