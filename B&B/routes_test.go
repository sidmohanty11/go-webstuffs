package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing

	default:
		t.Error(fmt.Sprintf("not chi mux, but is %T", v))
	}
}
